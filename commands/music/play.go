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
		vc, ok := ctx.VoiceContext.VCs[ctx.GuildID]
		if !ok {
			return fmt.Errorf("cannot play when im not connected")
		}

		queue, ok := ctx.VoiceContext.Queues[ctx.GuildID]
		if !ok {
			return fmt.Errorf("this isnt supposed to happen what")
		}

		videoUrl, err := ctx.GetArgIndex(0)
		if err != nil {
			return errors.New("please pass in a valid video URL")
		}

		queue.Push(voice.NewYTSrc(videoUrl, ctx.VoiceContext.Ytdl))
		//fp, _ := filepath.Abs("assets/test.mp3")
		//fmt.Println(fp)
		//queue.Push(&voice.StaticFile{FilePath: fp})
		if queue.Length() == 1 {
			song := queue.Pop()
			//ctx.Reply(fmt.Sprintf("%v", song))
			_, err := song.Play(vc)
			if err != nil {
				return fmt.Errorf("failed to play song: %v", err)
			}
			//err = <-doneChan
		}
		return nil
	},
}
