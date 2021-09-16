package av

import (
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util/av"
)

var Instakill = &framework.Command{
	Name: "instakill",
	Help: "Generates instakill audio on top of image (experimental)",
	Exec: func(ctx *framework.Context) error {
		img, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		videoBuffer := av.AddAudioToImage(img, "instakill.ogg")
		ctx.ReplyFile("instakill.webm", videoBuffer)
		return nil
	},
}
