/**
 *Created by XieJian on 2020/11/18 17:10
 *@Desc:
 */
package models

// 搜索关键词
type Search struct {
	SearchType string
	SearchURL  string
	Keywords   string
}

// 搜索内容
type SearchContent struct {
	Title   string
	Origin  string
	Content string
	IsParse bool
}

type Book struct {
	Title    string
	Describe string
	Avatar   string
}
