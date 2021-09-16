package av

import (
	"fmt"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/segment"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"image"
	"image/color"
	"strconv"
)

var Persona = &framework.Command{
	Name: "persona",
	Help: "Generates a persona style template (might be expanded in the future)",
	Exec: func(ctx *framework.Context) error {
		im, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		thr := ctx.GetArgIndexDefault(0, "90")
		thrInt, err := strconv.Atoi(thr)
		if err != nil {
			return fmt.Errorf("%s is not a number", thr)
		}
		im = effect.Grayscale(im)
		im = segment.Threshold(im, uint8(thrInt))
		im = effect.Median(im, 1.2)
		im = effect.Erode(im, .5)
		for y := 0; y < im.Bounds().Max.Y; y++ {
			for x := 0; x < im.Bounds().Max.X; x++ {
				r, g, b, _ := im.At(x, y).RGBA()
				if r<<8 >= 255 && g<<8 >= 255 && b<<8 >= 255 {
					im.(*image.RGBA).Set(x, y, color.RGBA{R: 255, G: 0, B: 0})
				}
			}
		}
		ctx.ReplyJPGImg(im, "persona.jpg")
		return nil
	},
}
