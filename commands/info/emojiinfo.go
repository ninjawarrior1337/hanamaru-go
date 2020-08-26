package info

import (
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
)

var EmojiInfo = &framework.Command{
	Name:               "einfo",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		msg, err := ctx.GetPreviousMessage()
		if err != nil {
			return err
		}
		emojis := util.ParseEmojis(msg.Content)
		emojiInfoMsg := ""
		for _, emoji := range emojis {
			emojiInfoMsg += fmt.Sprintf("%s: %s\n", emoji.Name, emoji.ID)
		}
		ctx.Reply(emojiInfoMsg)
		return nil
	},
	Setup: nil,
}
