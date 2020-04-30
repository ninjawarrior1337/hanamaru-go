package image

import (
	"bytes"
	"github.com/disintegration/imaging"
	"hanamaru/hanamaru"
	"image/png"
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
		var pngBuf = new(bytes.Buffer)
		png.Encode(pngBuf, resizedImg)
		ctx.ReplyFile("stretch.png", pngBuf)
		return nil
	},
}
