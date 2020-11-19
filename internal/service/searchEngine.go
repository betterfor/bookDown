/**
 *Created by XieJian on 2020/11/19 13:48
 *@Desc:
 */
package service

import (
	"fmt"
	"github.com/betterfor/bookDown/internal/models"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"log"
	"net/url"
	"strings"
)

const (
	defaultEngine    = "Baidu"
	defaultEngineURL = "https://www.baidu.com"
)

type SearchKeywords interface {
	Search(keywords string) ([]*models.SearchResult, error)
}

type SearchEngine struct {
	parseRule  string
	searchRule string
	domain     string
}

func NewBaiduSearchEngine() SearchKeywords {
	return &SearchEngine{
		parseRule:  "#content_left h3.t a",
		searchRule: "intitle:%s 小说 阅读",
		domain:     "http://www.baidu.com/s?wd=%s&ie=utf-8&rn=15&vf_bl=1", // rn条数;vf_bl页数
	}
}

// 搜索引擎查询数据
func (e *SearchEngine) Search(keywords string) ([]*models.SearchResult, error) {
	var searchRet []*models.SearchResult
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	searchKey := url.QueryEscape(fmt.Sprintf(e.searchRule, keywords))
	requestURL := fmt.Sprintf(e.domain, searchKey)

	c.OnHTML(e.parseRule, func(element *colly.HTMLElement) {
		var ret = &models.SearchResult{Title: element.Text, Href: element.Attr("href")}
		if checkURL(ret) {
			searchRet = append(searchRet, ret)
		}
	})

	err := c.Visit(requestURL)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// 根据搜索结果遍历有效性
func checkURL(result *models.SearchResult) bool {
	var valid bool
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnResponse(func(resp *colly.Response) {
		realURl := resp.Request.URL.String()
		fmt.Println(realURl)
		// 没有找到源地址/公司屏蔽网址
		if strings.Contains(realURl, "baidu") || strings.HasPrefix(realURl, "http://192.168.100.242") {
			return
		}

		// 过滤名单
		host := resp.Request.URL.Host

		result.Href = realURl
		result.IsParse = false
		result.Host = host
		valid = true
	})

	err := c.Visit(result.Href)
	if err != nil {
		log.Println("check ", result.Href, err)
	}
	return valid
}
