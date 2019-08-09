/**
 *Created by Xie Jian on 2019/8/9 13:29
 */
package utils

import "github.com/betterfor/gotils/colorString"

func Bold(message string) string {
	return colorString.Color("[bold]" + message)
}

func MagentaBold(message string) string {
	return Bold("[magenta]" + message)
}

func RedBold(message string) string {
	return Bold("[red]" + message)
}

func EndLine() string {
	return "\n"
}
