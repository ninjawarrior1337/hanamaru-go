package fun

import (
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

// Im bored lol
var Talk = &framework.Command{
	Name:               "talk",
	PermissionRequired: 0,
	OwnerOnly:          true,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		chatId, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		message, err := ctx.GetArgIndex(1)
		if err != nil {
			return err
		}
		_, err = ctx.Hanamaru.ChannelMessageSend(chatId, message)
		if err != nil {
			ctx.Reply("message failed to send, maybe the channel doesn't exist")
		}
		return nil
	},
}
