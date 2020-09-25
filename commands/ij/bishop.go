// +build ij

package ij

import (
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/markbates/pkger"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"golang.org/x/image/font"
	"image"
	"io/ioutil"
)

var baseImg image.Image
var fontFace font.Face

var Bishop = &framework.Command{
	Name: "bic",
	Exec: func(ctx *framework.Context) error {
		prevMsg, err := ctx.GetPreviousMessage()
		if err != nil {
			return err
		}
		madText := prevMsg.Content
		if madText == "" {
			return fmt.Errorf("please use this command after a message that contains text")
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
				textCtx.DrawStringWrapped(madText, x, y, 0.5, 0.5, float64(textCtx.Width()), 1, gg.AlignCenter)
			}
		}
		textCtx.SetRGB(1, 1, 1)
		textCtx.DrawStringWrapped(madText, float64(textCtx.Width()/2), float64(textCtx.Height()/2), 0.5, 0.5, float64(textCtx.Width()), 1, gg.AlignCenter)

		finalImg := gg.NewContextForImage(baseImg)
		finalImg.DrawImage(textCtx.Image(), 505, 124)

		ctx.ReplyJPGImg(finalImg.Image(), "bishop")

		return nil
	},
	Setup: func() error {
		f, err := pkger.Open("/assets/imgs/thefrick.png")
		if err != nil {
			return errors.New("failed to load bishop base image")
		}
		baseImg, _, err = image.Decode(f)
		if err != nil {
			return errors.New("failed to decode bishop base image")
		}

		fontF, err := pkger.Open("/assets/LeagueGothic.ttf")
		if err != nil {
			return errors.New("failed to load impact prarsedFont")
		}
		entireFile, _ := ioutil.ReadAll(fontF)
		prarsedFont, _ := truetype.Parse(entireFile)
		fontFace = truetype.NewFace(prarsedFont, &truetype.Options{Size: 64})
		return nil
	},
}
