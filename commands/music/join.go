package music

import (
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var Join = &framework.Command{
	Name:               "join",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		channel, err := ctx.GetSenderVoiceChannel()
		if err != nil {
			return err
		}

		err = ctx.VoiceContext.JoinChannel(ctx.Session, ctx.GuildID, channel.ID, ctx.ChannelID)
		if err != nil {
			return err
		}

		return nil
	},
}
