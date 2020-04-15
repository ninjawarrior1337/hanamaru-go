//+build jp

package jp

import (
	"bytes"
	"image/png"
	"io"
	"os"
	"testing"
)

func TestRenderRune(t *testing.T) {
	img := RenderRune("せ", 0, 1, 1)

	buff := new(bytes.Buffer)

	png.Encode(buff, img)
	file, _ := os.OpenFile("aaa.png", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	defer file.Close()

	io.Copy(file, buff)
}

func TestRenderPitchAccent(t *testing.T) {
	img, err := RenderPitchAccent("にほんじん", []int{0, 1, 1, 1, 0})
	if err != nil {
		t.Error(err)
	}

	buff := new(bytes.Buffer)

	png.Encode(buff, img)
	file, _ := os.OpenFile("bbb.png", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	defer file.Close()

	io.Copy(file, buff)
}
