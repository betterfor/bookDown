package site

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/betterfor/BookDown/models"
	"github.com/betterfor/BookDown/utils"
	"github.com/betterfor/GoLogger/logger"
	"github.com/betterfor/GoRequest"
	"github.com/betterfor/gotils/epubtil"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	biquge = "https://www.biqiuge.com/"
)

var (
	regx   = regexp.MustCompile(`\d{1,2}`)
	coder  = mahonia.NewDecoder("gbk") // gbk to utf8
	adRegx = regexp.MustCompile(`https:[\s\S]+wap.biqiuge.com`)
)

func Search(keyword string) (n int, slice models.SearchNovelSlice, err error) {
	url := "https://so.biqusoso.com/s.php?siteid=biqiuge.com&q=" + keyword
	resp, _, errs := GoRequest.New().Get(url).End()
	if len(errs) != 0 || errs != nil {
		msg := fmt.Sprintf("get http errors: %s", errs)
		logger.Error("[Get89HttpProxy] ", msg)
		err = errors.New(msg)
		return 0, nil, err
	}
	defer resp.Body.Close()
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	doc.Find(".search-list").Find("li").Each(func(i int, selection *goquery.Selection) {
		selection = selection.Find("span")
		var search models.Novel
		if regx.MatchString(selection.Eq(0).Text()) {
			search.NovelName = selection.Eq(1).Text()
			search.NovelURL, _ = selection.Eq(1).Find("a").Attr("href")
			search.Author = selection.Eq(2).Text()
			slice = append(slice, search)
		}
	})
	return len(slice), slice, nil
}

func Download(novel models.Novel) error {
	if novel.NovelURL == "" {
		return errors.New("no url to download")
	}

	chapterList, err := getChapters(&novel)
	if err != nil {
		return err
	}

	// get content
	Len := len(chapterList) - 6
	var ChanResp = make(chan models.ChanResp, Len)
	for _, chapter := range chapterList {
		if chapter.ChapterIndex < 0 {
			continue
		}
		logger.Infof("%+v\n", chapter)
		go GetContent(chapter, ChanResp)
	}

	var chapters = make([]models.Chapter, Len)
	for i := 0; i < Len; i++ {
		select {
		case v, ok := <-ChanResp:
			if !ok {
				continue
			} else {
				if v.Status == models.Failure {
					return errors.New(v.ErrMsg)
				} else {
					c := v.Data.(models.Chapter)
					chapters[c.ChapterIndex] = c
				}
			}
		}
	}
	novel.Chapters = chapters
	// 再写文件
	return writeFile(novel)
}

func GetContent(chapter models.Chapter, ch chan models.ChanResp) {
	var resp models.ChanResp
	// 需要使用ippool
	var times int
Retry:
	content, err := getContent(biquge + chapter.LinkURL)
	if err != nil && times < 100 {
		times++
		goto Retry
	} else if err != nil {
		logger.Errorf("[GetContent] Really can not download %+v, error:%s", chapter, err)
		resp.Status = models.Failure
		resp.ErrMsg = err.Error()
	} else {
		resp.Status = models.Success
		chapter.Content = content
		resp.Data = chapter
		ch <- resp
	}
	return
}

func getChapters(novel *models.Novel) ([]models.Chapter, error) {
	var chapterList []models.Chapter
	// get novel chapters
	resp, _, errs := GoRequest.New().Get(novel.NovelURL).AppendHeader("User-Agent", utils.RandomUA()).End()
	if len(errs) != 0 || errs != nil {
		msg := fmt.Sprintf("get http errors: %s", errs)
		logger.Error("[Get89HttpProxy] ", msg)
		err := errors.New(msg)
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	bookInfo := doc.Find(".book").Find(".info")
	image, exist := bookInfo.Find(".cover").Find("img").Attr("src")

	if exist {
		novel.Image = biquge + image
		novel.NovelName = coder.ConvertString(bookInfo.Find("h2").Text())
		novel.Author = coder.ConvertString(bookInfo.Find(".small").Find("span").Eq(0).Text())
		novel.Introduction = coder.ConvertString(bookInfo.Find(".intro").Text()) // todo 去掉后面的广告

		doc.Find(".listmain").Find("dl").Find("dd").Each(func(i int, selection *goquery.Selection) {
			//fmt.Println(enc.ConvertString(selection.Text()))

			url, exist := selection.Find("a").Attr("href")
			if exist {
				title := coder.ConvertString(selection.Find("a").Text())
				// 从-6开始是因为有6章最新章节
				var chapter = models.Chapter{ChapterIndex: i - 6, ChapterName: title, LinkURL: url}
				chapterList = append(chapterList, chapter)
			}
		})
		return chapterList, nil
	}
	return nil, errors.New("cannot find image")
}

func getContent(chapterURL string) ([]string, error) {
	ip := &models.Ippool{}
	ip, _ = ip.GetOneRandom()
	resp, _, errs := GoRequest.New().Get(chapterURL).AppendHeader("User-Agent", utils.RandomUA()).End()
	if len(errs) != 0 || errs != nil {
		msg := fmt.Sprintf("errors:%s", errs)
		return nil, errors.New(msg)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error("[DownloadNovelList] parse body error:", err)
		return nil, err
	}
	content := coder.ConvertString(doc.Find(".showtxt").Text())
	contents := biqiugeRexp(content)
	return contents, nil
}

func biqiugeRexp(content string) []string {
	content = strings.Replace(content, "聽", "", -1)
	content = strings.Replace(content, adRegx.FindString(content), "", 1)
	contents := strings.Split(content, "\n\n")
	return contents
}

func writeFile(novel models.Novel) error {
	// 一张图片
	var imagePath string
	url := biquge + novel.Image
	names := strings.Split(novel.Image, "/")
	imageName := names[len(names)-1]

	resp, err := http.Get(url)
	if err != nil {
		imagePath = ""
	} else {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			imagePath = ""
		} else {
			defer resp.Body.Close()

			imagePath = "./tmp/" + imageName
			image, err := os.Create(imagePath)
			if err != nil {
				imagePath = ""
			} else {
				defer os.Remove(imagePath)
				defer image.Close()
				image.Write(data)
			}
		}
	}
	var volume epubtil.Volume
	for _, chapter := range novel.Chapters {
		var cha = epubtil.Chapter{
			Name: chapter.ChapterName,
			Text: chapter.Content,
		}
		volume.Chapters = append(volume.Chapters, cha)
	}
	var epub = epubtil.ConEpub{
		BookName:      novel.NovelName,
		Author:        novel.Author,
		Description:   novel.Introduction,
		BookImagePath: imagePath,
		BookImageName: imageName,
		Volume:        []epubtil.Volume{volume},
	}
	return epubtil.ConvertEpub(epub, "./tmp/"+novel.NovelName+".epub")
}
