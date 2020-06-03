package music

import (
	"errors"
	"fmt"
	"hanamaru/hanamaru"
	"hanamaru/hanamaru/voice"
)

var Play = &hanamaru.Command{
	Name:               "play",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		_, ok := ctx.VoiceContext.VCs[ctx.GuildID]
		if !ok {
			return fmt.Errorf("cannot play when im not connected")
		}

		queueChan, ok := ctx.VoiceContext.QueueChannels[ctx.GuildID]
		if !ok {
			return fmt.Errorf("this isnt supposed to happen wot")
		}

		videoUrl, err := ctx.GetArgIndex(0)
		if err != nil {
			return errors.New("please pass in a valid video URL")
		}

		queueChan <- voice.NewYTSrc(videoUrl, ctx.VoiceContext.Ytdl)

		return nil
	},
}
