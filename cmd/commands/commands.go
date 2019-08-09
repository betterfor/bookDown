/**
 *Created by Xie Jian on 2019/8/9 10:45
 */
package commands

import (
	"flag"
	"github.com/betterfor/BookDown/utils"
	"io"
	"os"
	"strings"
)

type Command struct {
	// Run runs the command. The args are the arguments after the command name.
	Run func(cmd *Command, args []string) int

	// PreRun performs an operation before running the command
	PreRun func(cmd *Command, args []string)

	// UsageLine is one-line Usage message. The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'go help' output
	Short string

	// Long is the long message shown in the 'go help <command>' output
	Long string

	// Flag is a set of flags specific to this command
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own flag parsing.
	CustomFlags bool

	// output out writer if set in SetOutput(w)
	output *io.Writer
}

var AvailableCommands []*Command
var cmdUsage = `Use {{printf "bee help %s" .Name | bold}} for more information.{{endline}}`

// Name return the command's name: the first word in the Usage line
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ") // 找到subStr的位置
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// SetOutput sets the description for Usage and error messages. If output is nil, os.Stderr is uesd.
func (c *Command) SetOutput(output io.Writer) {
	c.output = &output
}

//
func (c *Command) Out() io.Writer {
	if c.output != nil {
		return *c.output
	}
	return os.Stderr
}

// Usage output the Usage for command.
func (c *Command) Usage() {
	utils.Tmpl(cmdUsage, c)
	os.Exit(2)
}

// Runnable reports whether the command can be run;
// otherwise it is a documention fake command such as import path
func (c *Command) Runnable() bool {
	return c.Run != nil
}

//
func (c *Command) Options() map[string]string {
	options := make(map[string]string)
	c.Flag.VisitAll(func(f *flag.Flag) {
		defaultVal := f.DefValue
		if len(defaultVal) > 0 {
			options[f.Name+"="+defaultVal] = f.Usage
		} else {
			options[f.Name] = f.Usage
		}
	})
	return options
}
