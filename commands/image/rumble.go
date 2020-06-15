package image

import (
	"errors"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/markbates/pkger"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"image"
)

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
		file, err := pkger.Open("/assets/imgs/rumble.png")
		if err != nil {
			return errors.New("failed to open rumble.png")
		}

		rumbleImg, _, err = image.Decode(file)
		if err != nil {
			return errors.New("failed to decode rumble.png")
		}
		return nil
	},
}
