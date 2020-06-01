package image

import (
	"bytes"
	"hanamaru/hanamaru"
	"image/jpeg"
)

var Jpg = &hanamaru.Command{
	Name:               "jpg",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
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
