package image

import (
	"github.com/anthonynsimon/bild/effect"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var EdgeDet = &framework.Command{
	Name:               "edge",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Performs edge detection on an image and outputs the result",
	Exec: func(ctx *framework.Context) error {
		im, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		im = effect.EdgeDetection(im, 1.0)
		ctx.ReplyJPGImg(im, "edge.jpg")
		return nil
	},
}
