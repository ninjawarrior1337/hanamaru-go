package debug

import (
	"fmt"
	"hanamaru/hanamaru"
)

var ListArgs = &hanamaru.Command{
	Name:               "listargs",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		ctx.ChannelMessageSend(ctx.ChannelID, fmt.Sprintf("%v", ctx.Args))
		return nil
	},
}
