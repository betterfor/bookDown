/**
 *Created by XieJian on 2020/11/19 14:03
 *@Desc:
 */
package models

// 搜索结果
type SearchResult struct {
	Href    string `json:"href"`
	Title   string `json:"title"`
	IsParse bool   `json:"is_parse"`
	Host    string `json:"host"`
}
