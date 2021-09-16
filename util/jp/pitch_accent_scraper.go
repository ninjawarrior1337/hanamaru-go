//go:build jp

package jp

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func generateFormData(in string) map[string][]byte {
	return map[string][]byte{
		"data[Phrasing][text]":       []byte(in),
		"data[Phrasing][curve]":      []byte("advanced"),
		"data[Phrasing][accent]":     []byte("advanced"),
		"data[Phrasing][analyze]":    []byte("true"),
		"data[Phrasing][estimation]": []byte("crf"),
	}
}

var regex = regexp.MustCompile("(\\[[\\d,]+\\])")

func ScrapePitchAccent(in string) (string, []int) {
	if in == "" {
		return "", nil
	}

	c := colly.NewCollector()

	var outSlice []int
	var outPhrase string

	c.OnHTML("#phrasing_main > div:nth-child(9) > div > script", func(element *colly.HTMLElement) {
		if element.Text != "" {
			matches := regex.FindAllStringSubmatch(element.Text, -1)
			if len(matches) > 0 {
				json.NewDecoder(bytes.NewReader([]byte(matches[0][1]))).Decode(&outSlice)
			}
		}
	})

	c.OnHTML("#phrasing_main > div:nth-child(9) > div > div.phrasing_text", func(element *colly.HTMLElement) {
		outPhrase = strings.TrimSpace(element.Text)
	})

	c.PostMultipart("http://www.gavo.t.u-tokyo.ac.jp/ojad/phrasing/index", generateFormData(in))

	return outPhrase, outSlice
}
