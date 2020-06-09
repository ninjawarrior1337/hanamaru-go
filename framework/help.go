package framework

import (
	"fmt"
	"strings"
)

var help = &Command{
	Name: "help",
	Help: "Displays this message",
	Exec: func(ctx *Context) error {
		output := ""

		for _, cmd := range ctx.Hanamaru.commands {
			if cmd.Help != "" {
				output += fmt.Sprintf("**%v**: %v\n", strings.Title(cmd.Name), cmd.Help)
			}
		}

		output += "\n"
		output += "Commands without documentation: "

		for _, cmd := range ctx.Hanamaru.commands {
			if cmd.Help == "" {
				output += fmt.Sprintf("%v, ", strings.Title(cmd.Name))
			}
		}
		ctx.Reply(output)
		return nil
	},
}
