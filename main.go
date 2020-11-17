package main

import (
	"github.com/betterfor/bookDown/internal/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bookDown"
	app.Version = "0.0.1"
	app.Commands = []*cli.Command{
		cmd.Web,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
