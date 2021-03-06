package util

import (
	"github.com/gocolly/colly/v2"
	"regexp"
	"strconv"
	"strings"
)

var prg = regexp.MustCompile(`(\d+) pages?`)
var trg = regexp.MustCompile(`([a-zA-Z ]+)`)

type NHentai struct {
	Title string

	Tags map[string][]string

	PageCount int
	CoverArt  string

	URL string
}

func ParseNhentai(digits int) (NHentai, error) {
	c := colly.NewCollector()
	n := NHentai{Tags: map[string][]string{}}

	c.OnHTML("#cover > a > img", func(element *colly.HTMLElement) {
		if s := element.Attr("data-src"); s != "" {
			n.CoverArt = s
		}
	})

	c.OnHTML("#info > h1", func(element *colly.HTMLElement) {
		if t := element.Text; t != "" {
			n.Title = t
		}
	})

	c.OnHTML("#tags", func(element *colly.HTMLElement) {
		element.ForEach("div:not(.hidden)", func(i int, element *colly.HTMLElement) {
			matches := trg.FindStringSubmatch(element.Text)
			if len(matches) == 0 {
				return
			}
			tagName := strings.TrimSuffix(matches[0], ":")
			if tagName == "Pages" || tagName == "Uploaded" {
				return
			}
			var tagContent []string
			element.ForEach("span > a", func(i int, element *colly.HTMLElement) {
				matches := trg.FindStringSubmatch(element.Text)
				if len(matches) == 0 {
					return
				}
				tagContent = append(tagContent, strings.Title(strings.TrimSuffix(matches[0], " ")))
			})
			n.Tags[tagName] = tagContent
		})
	})

	//Page number
	c.OnHTML("#tags > div:nth-child(8) > span > a > span", func(element *colly.HTMLElement) {
		n.PageCount, _ = strconv.Atoi(element.Text)
	})

	n.URL = "https://nhentai.net/g/" + strconv.Itoa(digits)

	err := c.Visit(n.URL)
	if err != nil {
		return n, err
	}

	return n, nil

}
