//go:build jp

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

func TestRenderPitchAccentConcurrent(t *testing.T) {
	img, err := RenderPitchAccentConcurrent("にほんじん", []int{0, 1, 1, 1, 0})
	if err != nil {
		t.Error(err)
	}

	buff := new(bytes.Buffer)

	png.Encode(buff, img)
	file, _ := os.OpenFile("bbb.png", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	defer file.Close()

	io.Copy(file, buff)
}

func benchmarkRenderPitchAccent(b *testing.B, phrase string, pitchData []int) {
	for n := 0; n < b.N; n++ {
		RenderPitchAccent(phrase, pitchData)
	}
}

func benchmarkRenderPitchAccentConcurrent(b *testing.B, phrase string, pitchData []int) {
	for n := 0; n < b.N; n++ {
		RenderPitchAccentConcurrent(phrase, pitchData)
	}
}

func BenchmarkRenderPitchAccentNihongo(b *testing.B) {
	benchmarkRenderPitchAccent(b, "にほんご", []int{0, 1, 1, 1})
}
func BenchmarkRenderPitchAccentNihongoConcurrent(b *testing.B) {
	benchmarkRenderPitchAccentConcurrent(b, "にほんご", []int{0, 1, 1, 1})
}
func BenchmarkRenderPitchAccentILikeJapanese(b *testing.B) {
	benchmarkRenderPitchAccent(b, "にほんごがすきです", []int{0, 1, 1, 1, 1, 0, 1, 0, 0})
}
func BenchmarkRenderPitchAccentILikeJapaneseConcurrent(b *testing.B) {
	benchmarkRenderPitchAccentConcurrent(b, "にほんごがすきです", []int{0, 1, 1, 1, 1, 0, 1, 0, 0})
}
