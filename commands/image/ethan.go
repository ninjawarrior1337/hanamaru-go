package image

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/markbates/pkger"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
	"log"
)

var ffData []byte

func init() {
	file, err := pkger.Open("/assets/noto.ttf")
	if err != nil {
		log.Fatalf("Failed to load noto font: %v", err)
	}
	entireFile, _ := ioutil.ReadAll(file)
	ffData = entireFile
}

func GetFont(size float64) font.Face {
	f, err := truetype.Parse(ffData)
	if err != nil {
		log.Fatalf("Failed to load noto font: %v", err)
	}
	return truetype.NewFace(f, &truetype.Options{Size: size})
}

var Ethan = &framework.Command{
	Name:               "ethan",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Generates an Ethan style meme",
	Exec: func(ctx *framework.Context) error {
		img, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		text, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}

		textCtx := gg.NewContext(img.Bounds().Max.X, img.Bounds().Max.Y/3)
		textCtx.SetColor(color.White)
		textCtx.DrawRectangle(0, 0, float64(textCtx.Width()), float64(textCtx.Height()))
		textCtx.Fill()
		textCtx.SetColor(color.Black)
		textCtx.SetFontFace(GetFont(16))
		textCtx.DrawStringWrapped(text, float64(textCtx.Width()/2), float64(textCtx.Height()/2), 0.5, 0.5, float64(textCtx.Width())*0.8, 1.0, gg.AlignCenter)

		ethanCtx := gg.NewContext(img.Bounds().Max.X, img.Bounds().Max.Y+textCtx.Height())
		ethanCtx.DrawImage(textCtx.Image(), 0, 0)
		ethanCtx.DrawImage(img, 0, textCtx.Height())
		ctx.ReplyPNGImg(ethanCtx.Image(), "ethan.png")
		return nil
	},
}
