package av

import (
	"bytes"
	"image"
	"image/draw"
	"image/gif"
	"sync"

	_ "embed"

	"github.com/fogleman/gg"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
)

//go:embed assets/spies.gif
var spiesGifBytes []byte

type editData struct {
	Width  int
	Height int
	Ws     image.Image
}

func computeEditData(i image.Image, text string) editData {
	//Compute whitespace image
	ws := gg.NewContext(i.Bounds().Dx(), i.Bounds().Dy()/3)

	ws.DrawRectangle(0, 0, float64(ws.Width()), float64(ws.Height()))
	ws.SetHexColor("#ffffff")
	ws.Fill()

	ws.SetHexColor("#000000")
	ws.SetFontFace(util.GetNotoBoldFont(48))
	ws.DrawStringWrapped("*"+text, float64(ws.Width()/2), float64(ws.Height()/2), 0.5, 0.5, float64(ws.Width()), 1.2, gg.AlignCenter)

	return editData{
		Ws:     ws.Image(),
		Width:  i.Bounds().Dx(),
		Height: i.Bounds().Dy() + i.Bounds().Dy()/3,
	}
}

func insertText(i image.PalettedImage, e editData) *image.Paletted {
	finalImg := gg.NewContext(e.Width, e.Height)

	finalImg.DrawImage(e.Ws, 0, 0)
	finalImg.DrawImage(i, 0, e.Ws.Bounds().Dy())

	pi := image.NewPaletted(finalImg.Image().Bounds(), i.(*image.Paletted).Palette)
	draw.Draw(pi, pi.Rect, finalImg.Image(), finalImg.Image().Bounds().Min, draw.Src)

	return pi
}

func getSpiesGif() *gif.GIF {
	g, _ := gif.DecodeAll(bytes.NewReader(spiesGifBytes))
	return g
}

var Wrong = &framework.Command{
	Name: "wrong",
	Help: "You used the wrong your haha",
	Exec: func(ctx *framework.Context) error {
		srcGif := getSpiesGif()
		text := ctx.TakeRest()
		ed := computeEditData(srcGif.Image[0], text)

		iArr := make([]*image.Paletted, len(srcGif.Image))
		var w sync.WaitGroup
		for idx, im := range srcGif.Image {
			w.Add(1)
			go func(idx int, im *image.Paletted) {
				iArr[idx] = insertText(im, ed)
				w.Done()
			}(idx, im)
		}
		w.Wait()

		srcGif.Config.Height = ed.Height
		srcGif.Image = iArr

		ctx.ReplyGIFImg(srcGif, "wrong")

		return nil
	},
}
