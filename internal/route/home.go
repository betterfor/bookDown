/**
 *Created by XieJian on 2020/11/17 16:45
 *@Desc:
 */
package route

import "github.com/gin-gonic/gin"

func HomePage(c *gin.Context) {
	c.HTML(200, "index", map[string]string{})
}
