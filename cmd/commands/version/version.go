/**
 *Created by Xie Jian on 2019/8/9 15:50
 */
package version

import (
	"flag"
	"fmt"
	"github.com/betterfor/BookDown/cmd/commands"
	"runtime"
)

var CmdVersion = &commands.Command{
	UsageLine: "version",
	Short:     "Prints the current version",
	Long:      `Prints the current version and Go version alongside the platform information.`,
	Run:       versionCmd,
}

var outputFormat string

const version = "v0.1.0"

func init() {
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	//fs.StringVar(&outputFormat,"o","","Set the output format. Either json or yaml.")
	CmdVersion.Flag = *fs
	commands.AvailableCommands = append(commands.AvailableCommands, CmdVersion)
}

func versionCmd(cmd *commands.Command, args []string) int {
	//cmd.Flag.Parse(args)
	fmt.Println("book version" + ":" + version + " " + runtime.GOOS + "/" + runtime.GOARCH)
	return 0
}
