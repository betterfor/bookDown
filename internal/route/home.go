/**
 *Created by XieJian on 2020/11/17 16:45
 *@Desc:
 */
package route

import (
	"github.com/betterfor/bookDown/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 首页
func HomePage(c *gin.Context) {
	c.Set("title", "home")
	c.HTML(200, "index", c.Keys)
}

// 搜索页
func SearchPage(c *gin.Context) {
	c.Set("title", "search")
	t := time.Now()

	keywords := c.Query("keywords")
	c.Set("keywords", keywords)

	e := service.NewBaiduSearchEngine()
	result, err := e.Search(keywords)
	if err != nil {
		log.Println("search error:", err)
	}

	c.Set("searchContents", result)
	c.Set("count", len(result))
	c.Set("time", time.Since(t).Seconds())
	c.HTML(200, "search/search", c.Keys)
}

func ChapterPage(c *gin.Context) {

}
