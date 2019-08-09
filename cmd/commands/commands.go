/**
 *Created by Xie Jian on 2019/8/9 10:45
 */
package commands

import (
	"flag"
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

//
func (c *Command) Usage() {

}
