package image

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/markbates/pkger"
	"hanamaru/hanamaru"
	"image"
	"log"
)

var rumbleImg image.Image

func init() {
	file, err := pkger.Open("/assets/imgs/rumble.png")
	if err != nil {
		log.Fatalf("Failed to open rumble.png: %v", err)
	}

	rumbleImg, _, err = image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to process rumble.png: %v", err)
	}
}

var Rumble = &hanamaru.Command{
	Name:               "rumble",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
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
}
