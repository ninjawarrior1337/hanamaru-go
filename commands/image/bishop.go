package image

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/markbates/pkger"
	"golang.org/x/image/font"
	"hanamaru/hanamaru"
	"image"
	"image/jpeg"
	"io/ioutil"
)

var baseImg image.Image
var fontFace font.Face

func init() {
	f, err := pkger.Open("/assets/thefrick.png")
	if err != nil {
		panic("failed to load bishop base image")
	}
	baseImg, _, err = image.Decode(f)
	if err != nil {
		panic("failed to decode bishop base image")
	}

	fontF, err := pkger.Open("/assets/impact.ttf")
	if err != nil {
		panic("failed to load impact prarsedFont")
	}
	entireFile, _ := ioutil.ReadAll(fontF)
	prarsedFont, _ := truetype.Parse(entireFile)
	fontFace = truetype.NewFace(prarsedFont, &truetype.Options{Size: 32})
}

var Bishop = &hanamaru.Command{
	Name: "bic",
	Exec: func(ctx *hanamaru.Context) error {
		prevMsg, err := ctx.GetPreviousMessage()
		if err != nil {
			return err
		}
		madText := prevMsg.Content
		if madText == "" {
			return fmt.Errorf("please use ths command after a message that contains text")
		}
		//Code borrowed from meme.go
		textCtx := gg.NewContext(244, 376)
		textCtx.SetRGBA(1, 1, 1, 0)
		textCtx.Clear()
		textCtx.SetFontFace(fontFace)
		textCtx.SetRGB(0, 0, 0)
		n := 3
		for dy := -n; dy <= n; dy++ {
			for dx := -n; dx <= n; dx++ {
				if dx*dx+dy*dy >= n*n {
					// give it rounded corners
					continue
				}
				x := float64(textCtx.Width()/2) + float64(dx)
				y := float64(textCtx.Height()/2) + float64(dy)
				textCtx.DrawStringAnchored(madText, x, y, 0.5, 0.5)
			}
		}
		textCtx.SetRGB(1, 1, 1)
		textCtx.DrawStringAnchored(madText, float64(textCtx.Width()/2), float64(textCtx.Height()/2), 0.5, 0.5)

		finalImg := gg.NewContextForImage(baseImg)
		finalImg.DrawImage(textCtx.Image(), 505, 124)
		jpgOut := new(bytes.Buffer)
		jpeg.Encode(jpgOut, finalImg.Image(), nil)
		ctx.ReplyFile("bishop.jpg", jpgOut)
		return nil
	},
}
