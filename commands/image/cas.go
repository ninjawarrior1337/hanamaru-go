package image

import (
	"bytes"
	"fmt"
	"hanamaru/hanamaru"
	"image"
	"image/draw"
	"image/jpeg"
	"strconv"
)
import "github.com/esimov/caire"

func NewProcessor(width, height int) *caire.Processor {
	return &caire.Processor{
		SobelThreshold: 10,
		BlurRadius:     1,
		NewWidth:       width,
		NewHeight:      height,
		Percentage:     false,
		Square:         false,
		Debug:          false,
		Scale:          false,
		FaceDetect:     false,
		FaceAngle:      0,
		Classifier:     "",
	}
}

var CAS = &hanamaru.Command{
	Name:               "cas",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		img, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		if len(ctx.Args) < 2 {
			return fmt.Errorf("not enoguh args passed")
		}
		nW, _ := strconv.Atoi(ctx.Args[0])
		nH, _ := strconv.Atoi(ctx.Args[1])
		p := NewProcessor(nW, nH)

		b := img.Image().Bounds()
		m := image.NewNRGBA(b)
		draw.Draw(m, b, img.Image(), b.Min, draw.Src)

		imgOut, err := p.Resize(m)
		if err != nil {
			return fmt.Errorf("failed to CAS image: %v", err)
		}

		outBuf := new(bytes.Buffer)
		jpeg.Encode(outBuf, imgOut, nil)
		ctx.ReplyFile("bruh.jpg", outBuf)
		return nil
	},
}
