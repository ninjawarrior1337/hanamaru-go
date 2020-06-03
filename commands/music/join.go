package music

import (
	"fmt"
	"hanamaru/hanamaru"
	"hanamaru/hanamaru/voice"
	"io"
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

		if _, ok := ctx.VoiceContext.QueueChannels[ctx.GuildID]; !ok {
			ctx.VoiceContext.QueueChannels[ctx.GuildID] = make(chan voice.Playable, 1024)
		}

		go func() {
			for p := range ctx.VoiceContext.QueueChannels[ctx.GuildID] {
				_, done, err := p.Play(vc)
				if err != nil {
					ctx.Reply("Failed to play song, skipping to next one.")
					continue
				}
				if err = <-done; err != io.EOF {
					ctx.Reply("Failed to play song, skipping to next one: " + err.Error())
				}
			}
		}()

		for {
			if vc.Ready {
				ctx.Reply("Joined!")
				break
			}
		}

		return nil
	},
}
