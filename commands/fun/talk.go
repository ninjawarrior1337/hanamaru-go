package fun

import (
	"fmt"
	"hanamaru/hanamaru"
)

//Im bored lol
var Talk = &hanamaru.Command{
	Name:               "talk",
	PermissionRequired: 0,
	OwnerOnly:          true,
	Help:               "",
	Exec: func(ctx *hanamaru.Context) error {
		chatId, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		message, err := ctx.GetArgIndex(1)
		if err != nil {
			return err
		}
		_, err = ctx.ChannelMessageSend(chatId, message)
		if err != nil {
			fmt.Errorf("message failed to send, maybe the channel doesn't exist")
		}
		return nil
	},
}
