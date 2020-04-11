package music

import (
	"fmt"
	"hanamaru/hanamaru"
	"hanamaru/hanamaru/voice"
	"path/filepath"
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

		//queue.Push(&voice.YoutubeSrc{YtUrl:"https://www.youtube.com/watch?v=XQsMmtC91b4"})
		fp, _ := filepath.Abs("assets/test.mp3")
		queue.Push(&voice.StaticFile{FilePath: fp})
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
