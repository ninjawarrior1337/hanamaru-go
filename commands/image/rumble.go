package image

import (
	"bytes"
	"errors"
	"image"

	_ "embed"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

//go:embed assets/rumble.png
var rumbleBytes []byte

var rumbleImg image.Image

var Rumble = &framework.Command{
	Name:               "rumble",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		input, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		mutRCtx := imaging.Resize(rumbleImg, input.Bounds().Max.X, input.Bounds().Max.Y/2, imaging.Lanczos)
		inputCtx := gg.NewContextForImage(input)
		inputCtx.DrawImage(mutRCtx, 0, 0)

		ctx.ReplyJPGImg(inputCtx.Image(), "rumble")
		return nil
	},
	Setup: func() error {
		var err error
		rumbleImg, _, err = image.Decode(bytes.NewReader(rumbleBytes))
		if err != nil {
			return errors.New("failed to decode rumble.png")
		}
		return nil
	},
}
