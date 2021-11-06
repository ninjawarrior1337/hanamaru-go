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

		for _, cmd := range ctx.Hanamaru.Commands {
			if cmd.Help != "" {
				output += fmt.Sprintf("**%v**: %v\n", cmd.Name, cmd.Help)
			}
		}

		output += "\n"
		output += "Commands without documentation: "
		commandNames := []string{}
		for _, cmd := range ctx.Hanamaru.Commands {
			if cmd.Help == "" {
				commandNames = append(commandNames, cmd.Name)
			}
		}
		output += strings.Join(commandNames, ", ")
		output += "\n"
		output += "\n"
		output += "Listeners: "
		listenerNames := []string{}
		{
			for _, ev := range ctx.Hanamaru.EventListeners {
				listenerNames = append(listenerNames, ev.Name)
			}
		}
		output += strings.Join(listenerNames, ", ")
		ctx.Reply(output)
		return nil
	},
}
