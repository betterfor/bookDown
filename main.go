package main

import (
	"fmt"
	"github.com/betterfor/BookDown/site"
)

func main() {
	//fmt.Println(site.Download(models.Novel{NovelURL: "https://www.biqiuge.com/book/4772/"}))
	list, err := site.GetRank()
	fmt.Println(list, err)
}
