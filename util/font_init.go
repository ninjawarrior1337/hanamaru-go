package util

import (
	"embed"
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed fonts/*
var fontsFS embed.FS

func GetNotoFont(size float64) font.Face {
	file, _ := fontsFS.Open("fonts/noto.ttf")
	entireFile, _ := ioutil.ReadAll(file)
	f, err := truetype.Parse(entireFile)
	if err != nil {
		log.Fatalf("Failed to load noto font: %v", err)
	}
	return truetype.NewFace(f, &truetype.Options{Size: size})
}

func GetNotoBoldFont(size float64) font.Face {
	file, _ := fontsFS.Open("fonts/notoBold.ttf")
	entireFile, _ := ioutil.ReadAll(file)
	f, err := truetype.Parse(entireFile)
	if err != nil {
		log.Fatalf("Failed to load notoBold font: %v", err)
	}
	return truetype.NewFace(f, &truetype.Options{Size: size})
}
