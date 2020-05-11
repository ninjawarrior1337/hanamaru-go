//+build jp

package jp

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/markbates/pkger"
	"golang.org/x/image/font"
	"image"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"unicode/utf8"
)

const KanaWidth = 75
const KanaHeight = 150

const TopOffset = 25
const BottomOffset = 80
const CenterOffset = KanaWidth / 2
const DotRadius = 8

var ffData []byte

func init() {
	file, err := pkger.Open("/assets/noto.ttf")
	if err != nil {
		log.Fatalf("Failed to load noto font: %v", err)
	}
	entireFile, _ := ioutil.ReadAll(file)
	ffData = entireFile
}

func GetFont() font.Face {
	f, err := truetype.Parse(ffData)
	if err != nil {
		log.Fatalf("Failed to load noto font: %v", err)
	}
	return truetype.NewFace(f, &truetype.Options{Size: 32})
}

func RenderPitchAccent(phrase string, pitchInfo []int) (image.Image, error) {
	if phrase == "" || len(pitchInfo) == 0 {
		return nil, fmt.Errorf("invalid input")
	}

	if utf8.RuneCountInString(phrase) != len(pitchInfo) {
		return nil, fmt.Errorf("please make sure the phrase is sent with the correct pitch info")
	}

	ctx := gg.NewContext(KanaWidth*utf8.RuneCountInString(phrase), KanaHeight)

	for i, mora := range strings.Split(phrase, "") {
		prevPitch := GetPitchAccentIndex(i-1, pitchInfo)
		nextPitch := GetPitchAccentIndex(i+1, pitchInfo)
		img := RenderRune(mora, pitchInfo[i], prevPitch, nextPitch)
		ctx.DrawImage(img, i*KanaWidth, 0)
	}

	return ctx.Image(), nil
}

func RenderPitchAccentConcurrent(phrase string, pitchInfo []int) (image.Image, error) {
	if phrase == "" || len(pitchInfo) == 0 {
		return nil, fmt.Errorf("invalid input")
	}

	if utf8.RuneCountInString(phrase) != len(pitchInfo) {
		return nil, fmt.Errorf("please make sure the phrase is sent with the correct pitch info")
	}

	var wg sync.WaitGroup
	ctx := gg.NewContext(KanaWidth*utf8.RuneCountInString(phrase), KanaHeight)

	chars := strings.Split(phrase, "")
	for i, mora := range chars {
		prevPitch := GetPitchAccentIndex(i-1, pitchInfo)
		nextPitch := GetPitchAccentIndex(i+1, pitchInfo)
		go func(mora string, pos int) {
			wg.Add(1)
			defer wg.Done()
			runeImg := RenderRune(mora, pitchInfo[pos], prevPitch, nextPitch)
			ctx.DrawImage(runeImg, pos*KanaWidth, 0)
		}(mora, i)
	}

	wg.Wait()
	return ctx.Image(), nil
}

func RenderRune(char string, pitchPlacement, pPitch, nPitch int) image.Image {
	ctx := gg.NewContext(KanaWidth, KanaHeight)

	ctx.SetFontFace(GetFont())
	ctx.SetRGB255(255, 255, 255)
	ctx.DrawStringAnchored(char, CenterOffset, 125, 0.5, 0.5)

	var CurrentOffset float64

	if pitchPlacement > 0 {
		CurrentOffset = TopOffset
	} else {
		CurrentOffset = BottomOffset
	}
	ctx.DrawCircle(CenterOffset, CurrentOffset, DotRadius)
	ctx.Fill()
	ctx.SetLineWidth(DotRadius / 2)

	if pPitch == 0 {
		ctx.DrawLine(CenterOffset, CurrentOffset, -KanaWidth+CenterOffset, BottomOffset)
	} else if pPitch > 0 {
		ctx.DrawLine(CenterOffset, CurrentOffset, -KanaWidth+CenterOffset, TopOffset)
	}

	if nPitch == 0 {
		ctx.DrawLine(CenterOffset, CurrentOffset, KanaWidth+CenterOffset, BottomOffset)
	} else if nPitch > 0 {
		ctx.DrawLine(CenterOffset, CurrentOffset, KanaWidth+CenterOffset, TopOffset)
	}

	ctx.Stroke()

	return ctx.Image()
}

func GetPitchAccentIndex(idx int, pitchInfoArray []int) int {
	if idx < 0 || idx > len(pitchInfoArray)-1 {
		return -1
	}
	return pitchInfoArray[idx]
}
