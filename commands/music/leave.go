package music

import (
	"hanamaru/hanamaru"
)

var Leave = &hanamaru.Command{
	Name:               "leave",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		channel, err := ctx.GetSenderVoiceChannel()
		if err != nil {
			return err
		}
		err = ctx.VoiceContext.LeaveChannel(ctx.GuildID, channel.ID)
		if err != nil {
			return err
		}
		return nil
	},
}
