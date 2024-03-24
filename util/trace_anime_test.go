package util

import (
	"image"
	"net/http"
	"testing"
)

func loadImage(url string) (i image.Image) {
	resp, err := http.Get(url)
	if err != nil {
		panic("failed to download image")
	}
	if resp.StatusCode != 200 {
		panic("failed to download image")
	}
	i, _, _ = image.Decode(resp.Body)
	return i
}

func TestTraceAnime(t *testing.T) {
	SkipCI(t)
	i := loadImage("https://theglorioblog.files.wordpress.com/2016/07/theresistance.png?w=700&h=394")
	ta, err := TraceAnime(i)
	if err != nil {
		t.Error(err)
	}
	if ta.Result[0].Anilist != 21584 && ta.Result[0].Episode != 4 {
		t.Fail()
	}
}
