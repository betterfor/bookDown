package site

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/betterfor/BookDown/models"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/**
主要获取起点的排行榜和搜索结果
用来换源
*/
const (
	qidianRank   = "https://www.qidian.com/rank"
	qidianBook   = "https://book.qidian.com/info/"
	qidianSearch = "https://www.qidian.com/search"
)

// 起点排行榜
func GetRank() (ranks []models.RankBook, err error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var f *os.File

	err = chromedp.Run(ctx, visitWeb(qidianRank, "#login-btn", "div.rank-body"))
	if err != nil {
		return nil, err
	}
	fileName := "qidian_rank.txt"
	err = ioutil.WriteFile(fileName, []byte(res), os.ModePerm)
	if err != nil {
		return nil, err
	}
	f, err = os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer os.Remove(fileName)
	defer f.Close()
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(f)
	if err != nil {
		return nil, err
	}
	// 其中有个推荐月票排行榜，分周/月/总共30个
	doc.Find(".rank-list").Each(func(i int, selection *goquery.Selection) {
		var rank models.RankBook
		selection.Find(".wrap-title").Each(func(i int, selection *goquery.Selection) {
			rank.RankTitle = strings.TrimSpace(selection.Text())
		})
		selection.Find(".book-list").Find("li").Each(func(i int, sel *goquery.Selection) {
			var bookinfo models.BookInfo
			bookinfo.BookName = strings.TrimSpace(sel.Find("h4").Text())
			url, exist := sel.Find("h4").Find("a").Attr("href")
			if exist {
				bookinfo.BookURL = url
				rank.BookInfos = append(rank.BookInfos, bookinfo)
			}
			bookinfo.BookName = strings.TrimSpace(sel.Find(".name-box").Find("a").Text())
			url, exist = sel.Find(".name-box").Find("a").Attr("href")
			if exist {
				bookinfo.BookURL = strings.TrimSpace(url)
				rank.BookInfos = append(rank.BookInfos, bookinfo)
			}
		})
		ranks = append(ranks, rank)
	})
	return ranks, nil
}

var res string

func visitWeb(url, visible, target string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(visible, chromedp.ByQuery),
		chromedp.OuterHTML(target, &res),
	}
}

func GetSearch(keyword string, page int) (infos []models.BookInfo, err error) {
	url := qidianSearch + "?kw=" + keyword + "&page=" + strconv.Itoa(page)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var f *os.File
	err = chromedp.Run(ctx, visitWeb(url, "#pin-nav", "div.book-img-text"))
	if err != nil {
		return nil, err
	}
	fileName := "qidian_search.txt"
	err = ioutil.WriteFile(fileName, []byte(res), os.ModePerm)
	if err != nil {
		return nil, err
	}
	f, err = os.Open(fileName)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return nil, err
	}
	doc.Find(".res-book-item").Find(".book-mid-info").Each(func(i int, selection *goquery.Selection) {
		var info models.BookInfo
		info.BookName = selection.Find("h4").Find("a").Text()
		url, exist := selection.Find("h4").Find("a").Attr("href")
		if exist {
			info.BookURL = url
			infos = append(infos, info)
		}
	})
	return infos, err
}
