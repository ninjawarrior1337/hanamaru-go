package av

import (
	"bytes"
	"math/rand"

	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util/av"
)

var Instakill = &framework.Command{
	Name: "instakill",
	Help: "Generates instakill audio on top of image (experimental)",
	Exec: func(ctx *framework.Context) error {
		var useBaikenAudio = false
		var videoBuffer *bytes.Buffer
		if rand.Intn(22) == 21 {
			useBaikenAudio = true
		}
		img, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		baiken, _ := ctx.GetArgIndex(0)
		if baiken != "" || useBaikenAudio {
			videoBuffer = av.AddAudioToImage(img, "instakill_baiken.ogg")
		} else {
			videoBuffer = av.AddAudioToImage(img, "instakill.ogg")
		}
		ctx.ReplyFile("instakill.webm", videoBuffer)
		return nil
	},
}
