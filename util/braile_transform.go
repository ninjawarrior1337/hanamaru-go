package util

import (
	"github.com/fogleman/gg"
	"image"
	"image/draw"
)

var charList = `$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,"^'`

func TransformBraile(ctx *gg.Context) string {
	grayImg := image.NewGray(ctx.Image().Bounds())
	draw.Draw(grayImg, ctx.Image().Bounds(), ctx.Image(), ctx.Image().Bounds().Min, draw.Src)

	newCtx := gg.NewContextForImage(grayImg)
	newCtx.Scale(200, 100)

	b := newCtx.Image().Bounds()

	//var fString = ""

	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			//brightness := grayImg.At(x, y).
			//newIdx :=
			//fString+=
		}
	}
}
