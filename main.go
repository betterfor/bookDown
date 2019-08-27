package main

import (
	"flag"
	"github.com/betterfor/BookDown/cmd"
	"github.com/betterfor/BookDown/cmd/commands"
	"os"
)

func main() {
	flag.Usage = cmd.Usage

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		cmd.Usage()
		os.Exit(2)
		return
	}
	if args[0] == "help" {
		cmd.Help(args)
		return
	}
	for _, c := range commands.AvailableCommands {
		if c.Name() == args[0] && c.Run != nil {
			c.Flag.Usage = func() {
				c.Usage()
			}
		} else {
			c.Flag.Parse(args[1:])
			args = c.Flag.Args()
		}

		if c.PreRun != nil {
			c.PreRun(c, args)
		}

		os.Exit(c.Run(c, args))
		return
	}
}
