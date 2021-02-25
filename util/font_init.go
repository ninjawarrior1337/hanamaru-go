package util

import (
	"embed"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed fonts/*
var fontsFS embed.FS

func GetFontByName(name string, size float64) font.Face {
	file, _ := fontsFS.Open(fmt.Sprintf("fonts/%v.ttf", name))
	entireFile, _ := ioutil.ReadAll(file)
	f, err := truetype.Parse(entireFile)
	if err != nil {
		log.Fatalf("Failed to load font %v: %v", name, err)
	}
	return truetype.NewFace(f, &truetype.Options{Size: size})
}

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
