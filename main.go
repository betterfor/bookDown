package main

import (
	"fmt"
	"github.com/betterfor/BookDown/models"
	"github.com/betterfor/BookDown/site"
)

func main() {
	fmt.Println(site.Download(models.Novel{NovelURL: "https://www.biqiuge.com/book/4772/"}))
}
