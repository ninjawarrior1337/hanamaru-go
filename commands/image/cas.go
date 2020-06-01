package image

import (
	"fmt"
	"hanamaru/hanamaru"
	"image"
	"image/draw"
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
		widthArg, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		heightArg, err := ctx.GetArgIndex(1)
		if err != nil {
			return err
		}
		nW, _ := strconv.Atoi(widthArg)
		nH, _ := strconv.Atoi(heightArg)
		p := NewProcessor(nW, nH)

		b := img.Bounds()
		m := image.NewNRGBA(b)
		draw.Draw(m, b, img, b.Min, draw.Src)

		imgOut, err := p.Resize(m)
		if err != nil {
			return fmt.Errorf("failed to CAS image: %v", err)
		}

		ctx.ReplyJPGImg(imgOut, "cas")
		return nil
	},
}
