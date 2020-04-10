package music

import (
	"fmt"
	"hanamaru/hanamaru"
	"hanamaru/hanamaru/voice"
)

var Join = &hanamaru.Command{
	Name:               "join",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		channel, err := ctx.GetVoiceChannnel()
		if err != nil {
			return err
		}
		vc, err := ctx.ChannelVoiceJoin(ctx.GuildID, channel.ID, false, false)
		if err != nil {
			return fmt.Errorf("failed to join VC: %v", err)
		}
		ctx.VoiceContext.VCs[ctx.GuildID] = vc

		if _, ok := ctx.VoiceContext.Queues[ctx.GuildID]; !ok {
			ctx.VoiceContext.Queues[ctx.GuildID] = &voice.Queue{}
		}

		return nil
	},
}
