package music

import "github.com/ninjawarrior1337/hanamaru-go/framework"

var Leave = &framework.Command{
	Name:               "leave",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		channel, err := ctx.GetSenderVoiceChannel()
		if err != nil {
			return err
		}
		err = ctx.Hanamaru.VoiceContext.LeaveChannel(ctx.GuildID, channel.ID)
		if err != nil {
			return err
		}
		return nil
	},
}
