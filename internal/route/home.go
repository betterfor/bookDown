/**
 *Created by XieJian on 2020/11/17 16:45
 *@Desc:
 */
package route

import "github.com/gin-gonic/gin"

// 首页
func HomePage(c *gin.Context) {
	c.Set("title", "home")
	c.HTML(200, "index", c.Keys)
}

func SearchPage(c *gin.Context) {
	c.Set("title", "search")
	c.HTML(200, "search/search", c.Keys)
}
