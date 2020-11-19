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

type Book struct {
	Title    string
	Describe string
	Avatar   string
}
