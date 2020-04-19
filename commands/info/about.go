package info

import (
	"hanamaru/hanamaru"
)

var About = &hanamaru.Command{
	Name: "about",
	Exec: func(ctx *hanamaru.Context) error {
		_, _ = ctx.ChannelMessageSend(ctx.ChannelID, "bruh")
		return nil
	},
	OwnerOnly: true,
}
