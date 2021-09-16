//go:build jp

package jp

import (
	"bytes"

	"github.com/ninjawarrior1337/hanamaru-go/framework"

	"image/png"

	"github.com/ninjawarrior1337/hanamaru-go/util/jp"
)

var Pitch = &framework.Command{
	Name:               "pitch",
	Help:               "Generates a Dogen-style pitch accent diagram. Uses OJAD Suzuki-kun for pitch info",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		arg, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		phrase, pitchData := jp.ScrapePitchAccent(arg)
		img, err := jp.RenderPitchAccentConcurrent(phrase, pitchData)
		if err != nil {
			return err
		}
		buffer := new(bytes.Buffer)
		png.Encode(buffer, img)
		ctx.ReplyFile("pitch.png", buffer)
		return nil
	},
}
