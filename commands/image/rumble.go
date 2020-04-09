package image

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/markbates/pkger"
	"hanamaru/hanamaru"
	"image"
	"image/jpeg"
	"log"
)

var rumbleCtx *gg.Context

func init() {
	file, err := pkger.Open("/assets/rumble.png")
	if err != nil {
		log.Fatalf("Failed to open rumble.png: %v", err)
	}

	rumblePng, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to process rumble.png: %v", err)
	}

	rumbleCtx = gg.NewContextForImage(rumblePng)
}

var Rumble = &hanamaru.Command{
	Name:               "rumble",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		input, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		mutRCtx := gg.NewContextForImage(rumbleCtx.Image())
		mutRCtx.Scale(float64(input.Width()), float64(input.Height()/2))
		input.DrawImage(mutRCtx.Image(), 0, 0)

		buf := new(bytes.Buffer)
		jpeg.Encode(buf, input.Image(), nil)
		ctx.ChannelFileSend(ctx.ChannelID, "bruh.jpg", buf)
		return nil
	},
}
