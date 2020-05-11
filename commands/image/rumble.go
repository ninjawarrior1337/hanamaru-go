package image

import (
	"github.com/disintegration/imaging"
	"github.com/markbates/pkger"
	"hanamaru/hanamaru"
	"image"
	"log"
)

var rumbleImg image.Image

func init() {
	file, err := pkger.Open("/assets/rumble.png")
	if err != nil {
		log.Fatalf("Failed to open rumble.png: %v", err)
	}

	rumbleImg, _, err = image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to process rumble.png: %v", err)
	}
}

var Rumble *hanamaru.Command = &hanamaru.Command{
	Name:               "rumble",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		input, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		mutRCtx := imaging.Resize(rumbleImg, input.Width(), input.Height()/2, imaging.Lanczos)
		input.DrawImage(mutRCtx, 0, 0)

		ctx.ReplyJPGImg(input.Image(), "rumble")
		return nil
	},
}
