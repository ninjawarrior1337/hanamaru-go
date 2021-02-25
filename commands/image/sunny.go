package image

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
)

var Sunny = &framework.Command{
	Name: "sunny",
	Help: "Generates always sunny in philadelphia title cards",
	Exec: func(ctx *framework.Context) error {
		arg := ctx.TakeRest()
		arg = arg[1:]
		board := gg.NewContext(1920, 1080)
		board.SetColor(color.Black)
		board.Fill()
		board.SetColor(color.White)
		board.SetFontFace(util.GetFontByName("Textile", 128))
		board.DrawStringAnchored(fmt.Sprintf(`"%v"`, arg), float64(board.Width()/2), float64(board.Height()/2), 0.5, 0.5)

		ctx.ReplyJPGImg(board.Image(), "phily")

		return nil
	},
}
