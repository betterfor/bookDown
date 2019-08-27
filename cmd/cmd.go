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
{{range .}}{{if not .Runnable}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use {{"book help [topic]" | bold}} for more information about that topic.
`

var htlpTemplate = `{{"USAGE" | headline}}
  {{.UsageLine | printf "bee %s" | bold}}
{{if .Options}}{{endline}}{{"OPTIONS | headline}}{{range $k,$v := .Options}}
  {{$k | printf "-%s" | bold}}
      {{$v}}
  {{end}}{{end}}
{{"DESCRIPTION" | headline}}
  {{tmpltostr .Long . | trim}}"`

var ErrorTemplate = `book: %s.
Use {{"book help" | bold}} for more information`

func Usage() {
	utils.Tmpl(usageTemplate, commands.AvailableCommands)
}

func Help(args []string) {
	if len(args) == 0 {
		Usage()
	}
	if len(args) != 1 {
		utils.PrintErrorAndExit("Too many arguments.", ErrorTemplate)
	}

	arg := args[0]
	for _, cmd := range commands.AvailableCommands {
		if cmd.Name() == arg {
			utils.Tmpl(htlpTemplate, cmd)
			return
		}
	}
	utils.PrintErrorAndExit("Unknown help topic", ErrorTemplate)
}
