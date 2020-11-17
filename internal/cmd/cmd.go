/**
 *Created by XieJian on 2020/11/17 14:18
 *@Desc:
 */
package cmd

import (
	"github.com/betterfor/bookDown/internal/route"
	"github.com/betterfor/bookDown/internal/tpl"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var Web = &cli.Command{
	Name:        "web",
	Usage:       "start web server",
	Description: "display novel list,novel content and novel download page with ui",
	Action:      runWeb,
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "port,p", Usage: "expose port to serve (default 3000)", Value: "3000"},
	},
}

func runWeb(c *cli.Context) error {
	g := newGinMiddle()
	t := tpl.LoadTemplate("templates/")
	g.SetHTMLTemplate(t)

	g.Static("/css", "./public/css")
	g.Static("/fonts", "./public/fonts")
	g.Static("/img", "./public/img")
	g.Static("/js", "./public/js")

	g.GET("/", route.HomePage)

	return g.Run(":" + c.String("port"))
}

func newGinMiddle() *gin.Engine {
	g := gin.Default()

	//t,err := loadTpl()
	//if err != nil {
	//	panic(err)
	//}
	g.FuncMap = tpl.FuncMap()

	return g
}

// todo use golang 1.16 embed to package files
//func loadTpl() (*template.Template, error) {
//	t := template.New("")
//	t.Funcs(tpl.FuncMap())
//}
