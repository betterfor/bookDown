/**
 *Created by Xie Jian on 2019/8/9 16:27
 */
package cmd

import (
	"github.com/betterfor/BookDown/cmd/commands"
	_ "github.com/betterfor/BookDown/cmd/commands/version"
	"github.com/betterfor/BookDown/utils"
)

var usageTemplate = `BookDown is a useful tool to download novels.

{{"USAGE" | headline}}
	{{"book command [arguments]" | bold}}
{{"AVAILABLE COMMANDS" | headline}}
{{range .}}{{if .Runnable}}
	{{.Name | printf "%-11s" | bold}} {{.Short}}{{end}}{{end}}

Use {{"book help [command]" | bold}} for more information about a command.

{{"ADDITIONAL HELP TOPICS" | headline}}
{{range .}}{{if not .Runable}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use {{"book help [topic]" | bold}} for more information about that topic.
`

func Usage() {
	utils.Tmpl(usageTemplate, commands.AvailableCommands)
}
