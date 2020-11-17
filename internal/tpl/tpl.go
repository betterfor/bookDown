/**
 *Created by XieJian on 2020/11/17 15:27
 *@Desc:
 */
package tpl

import (
	"html/template"
	"sync"
)

var (
	funcMap     template.FuncMap
	funcMapOnce sync.Once
)

func FuncMap() template.FuncMap {
	funcMapOnce.Do(func() {

	})
	return funcMap
}
