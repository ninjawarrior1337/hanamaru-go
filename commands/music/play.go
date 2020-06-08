package music

import (
	"errors"
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/framework/voice"
)

var Play = &framework.Command{
	Name:               "play",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
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
