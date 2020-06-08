package image

import (
	"bytes"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"image/jpeg"
)

var Jpg = &framework.Command{
	Name:               "jpg",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		img, err := ctx.GetImage(0)
		if err != nil {
			return err
		}

		outBuf := new(bytes.Buffer)
		err = jpeg.Encode(outBuf, img, &jpeg.Options{Quality: 1})
		if err != nil {
			return err
		}
		ctx.ReplyFile("bruh.jpg", outBuf)
		return nil
	},
}
