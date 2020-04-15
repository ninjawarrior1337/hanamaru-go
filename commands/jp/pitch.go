// +build jp

package jp

import (
	"bytes"
	"hanamaru/hanamaru"
	"hanamaru/util/jp"
	"image/png"
)

var Pitch = &hanamaru.Command{
	Name:               "pitch",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		arg, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		phrase, pitchData := jp.ScrapePitchAccent(arg)
		img, err := jp.RenderPitchAccent(phrase, pitchData)
		if err != nil {
			return err
		}
		buffer := new(bytes.Buffer)
		png.Encode(buffer, img)
		ctx.ReplyFile("pitch.png", buffer)
		return nil
	},
}
