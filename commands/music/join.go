package music

import (
	"hanamaru/hanamaru"
)

var Join = &hanamaru.Command{
	Name:               "join",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
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
