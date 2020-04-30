package image

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/markbates/pkger"
	"hanamaru/hanamaru"
	"image"
	"image/jpeg"
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

var Rumble = &hanamaru.Command{
	Name:               "rumble",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		input, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		mutRCtx := imaging.Resize(rumbleImg, input.Width(), input.Height()/2, imaging.Lanczos)
		input.DrawImage(mutRCtx, 0, 0)

		buf := new(bytes.Buffer)
		jpeg.Encode(buf, input.Image(), nil)
		ctx.ChannelFileSend(ctx.ChannelID, "bruh.jpg", buf)
		return nil
	},
}
