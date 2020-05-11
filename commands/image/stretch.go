package image

import (
	"github.com/disintegration/imaging"
	"hanamaru/hanamaru"
)

var Stretch = &hanamaru.Command{
	Name:               "stretch",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Exec: func(ctx *hanamaru.Context) error {
		img, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		resizedImg := imaging.Resize(img.Image(), img.Width()*2, img.Height(), imaging.Lanczos)
		ctx.ReplyPNGImg(resizedImg, "resize")
		return nil
	},
}
