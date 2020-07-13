package image

import (
	"fmt"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/segment"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
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
		ctx.ReplyJPGImg(im, "persona.jpg")
		return nil
	},
}
