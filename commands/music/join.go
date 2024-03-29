package music

import (
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var Join = &framework.Command{
	Name:               "join",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		channel, err := ctx.GetSenderVoiceChannel()
		optionalVCId := ctx.GetArgIndexDefault(0, "")

		if len(optionalVCId) > 0 {
			err = ctx.Hanamaru.VoiceContext.JoinChannel(ctx.Hanamaru.Session, ctx.GuildID, optionalVCId, ctx.ChannelID)
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		err = ctx.Hanamaru.VoiceContext.JoinChannel(ctx.Hanamaru.Session, ctx.GuildID, channel.ID, ctx.ChannelID)
		if err != nil {
			return err
		}

		return nil
	},
}
