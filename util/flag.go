package util

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/fogleman/gg"
	"image"
)

const WIDTH = 1920
const HEIGHT = 1080

func EncodeFlag(data []byte) image.Image {
	hexString := hex.EncodeToString(data)

	//Pad hex string with leading zeroes
	for i := 0; i < len(hexString)%6; i++ {
		hexString = "0" + hexString
	}
	//Setup context
	ctx := gg.NewContext(WIDTH, HEIGHT)
	//Compute width
	colWidth := WIDTH / (len(hexString) / 6)

	for i := 0; i < len(hexString)/6; i += 1 {
		currColor := hexString[6*i : (6*i)+6]
		ctx.SetHexColor(currColor)
		ctx.DrawRectangle(float64(i*colWidth), 0, float64(colWidth), HEIGHT)
		ctx.Fill()
	}
	return ctx.Image()
}

func rgbToHex(r, g, b uint32) string {
	return fmt.Sprintf("%02x%02x%02x", r, g, b)
}

func DecodeFlag(image image.Image) ([]byte, error) {
	hexString := ""
	prevHex := ""
	for x := 0; x < image.Bounds().Dx(); x++ {
		r, g, b, _ := image.At(x, 0).RGBA()
		currHex := rgbToHex(r>>8, g>>8, b>>8)
		if currHex != prevHex {
			hexString = hexString + currHex
			prevHex = currHex
		}
	}
	data, err := hex.DecodeString(hexString)
	data = bytes.Trim(data, "\x00")
	return data, err
}
