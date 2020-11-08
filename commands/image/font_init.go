package image

import (
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/markbates/pkger"
	"golang.org/x/image/font"
)

func GetNotoFont(size float64) font.Face {
	file, _ := pkger.Open("/assets/noto.ttf")
	entireFile, _ := ioutil.ReadAll(file)
	f, err := truetype.Parse(entireFile)
	if err != nil {
		log.Fatalf("Failed to load noto font: %v", err)
	}
	return truetype.NewFace(f, &truetype.Options{Size: size})
}

func GetNotoBoldFont(size float64) font.Face {
	file, _ := pkger.Open("/assets/notoBold.ttf")
	entireFile, _ := ioutil.ReadAll(file)
	f, err := truetype.Parse(entireFile)
	if err != nil {
		log.Fatalf("Failed to load notoBold font: %v", err)
	}
	return truetype.NewFace(f, &truetype.Options{Size: size})
}
