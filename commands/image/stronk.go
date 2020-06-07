package image

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/markbates/pkger"
	"golang.org/x/image/colornames"
	"hanamaru/hanamaru"
	"image"
)

var Stronk = &hanamaru.Command{
	Name:               "stronk",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Generates a stronk image, I made this out of spite: <text> [hex color]",
	Exec: func(ctx *hanamaru.Context) error {
		text, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		color := ctx.GetArgIndexDefault(1, "#ff006e")

		fin := gg.NewContextForImage(getStronkImage())
		w, h := fin.MeasureString(text)

		textCtx := gg.NewContext(int(w), int(h))
		textCtx.SetHexColor(color)
		textCtx.DrawRectangle(0, 0, w, h)
		textCtx.Fill()
		textCtx.SetColor(colornames.Black)
		textCtx.DrawStringAnchored(text, w/2, h/2, 0.5, 0.5)
		textImg := imaging.Resize(textCtx.Image(), 142, 39, imaging.Lanczos)

		fin.DrawImage(textImg, 403, 284)
		ctx.ReplyJPGImg(fin.Image(), "stronk")
		return nil
	},
}

func getStronkImage() image.Image {
	f, _ := pkger.Open("/assets/imgs/stronk.png")
	i, _, _ := image.Decode(f)
	return i
}
