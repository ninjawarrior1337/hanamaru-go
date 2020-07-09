package jisho

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strconv"
	"strings"
)

type JLPTLevel string

const (
	N5 JLPTLevel = "N5"
	N4 JLPTLevel = "N4"
	N3 JLPTLevel = "N3"
	N2 JLPTLevel = "N2"
	N1 JLPTLevel = "N1"
)

var exampleRegex = regexp.MustCompile(`(\S+)\s+【(\S+)】\s+(.+)`)

//var radicalRegex = regexp.MustCompile(`(.+)\s(.+)\s\((.+)\)`)

type ReadingExample struct {
	Example string
	Reading string
	Meaning string
}

type Radical struct {
	Meaning string
	Symbol  string
}

type Kanji struct {
	found              bool
	Rune               rune
	TaughtIn           string
	JLPT               JLPTLevel
	NewspaperFreqRank  int
	StrokeCount        int
	Meaning            string
	Kunyomi            []string
	KunyomiExamples    []*ReadingExample
	Onyomi             []string
	OnyomiExamples     []*ReadingExample
	Radical            *Radical
	Parts              []string
	StrokeOrderDiagram string
	JishoUri           string
}

func parseExample(exampleRaw string) *ReadingExample {
	res := exampleRegex.FindAllStringSubmatch(exampleRaw, -1)
	if len(res) > 0 {
		var re ReadingExample
		re.Example = res[0][1]
		re.Reading = res[0][2]
		re.Meaning = res[0][3]
		return &re
	}
	return nil
}

func (j *Jisho) SearchForKanji(kanji string) (*Kanji, error) {
	rel := "/" + kanji + "%20%23kanji"
	u := j.BaseScrapeUri + rel
	c := colly.NewCollector()

	var k Kanji

	//Found
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-2.columns > div > div:nth-child(1) > h1", func(element *colly.HTMLElement) {
		if element.Text == "" {
			k.found = false
		} else {
			k.found = true
		}
	})

	//Rune
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-2.columns > div > div:nth-child(1) > h1", func(element *colly.HTMLElement) {
		k.Rune = []rune(element.Text)[0]
	})

	//Taught In
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-10.columns > div > div.small-12.large-5.columns > div > div.grade > strong", func(element *colly.HTMLElement) {
		k.TaughtIn = element.Text
	})

	//JLPT
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-10.columns > div > div.small-12.large-5.columns > div > div.jlpt > strong", func(element *colly.HTMLElement) {
		k.JLPT = JLPTLevel(element.Text)
	})

	//Newspaper Freq
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-10.columns > div > div.small-12.large-5.columns > div > div.frequency > strong", func(element *colly.HTMLElement) {
		k.NewspaperFreqRank, _ = strconv.Atoi(element.Text)
	})

	//Stroke Count
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-2.columns > div > div:nth-child(2) > div.kanji-details__stroke_count > strong", func(element *colly.HTMLElement) {
		k.StrokeCount, _ = strconv.Atoi(element.Text)
	})

	//Meaning
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-10.columns > div > div.small-12.large-7.columns.kanji-details__main > div.kanji-details__main-meanings", func(element *colly.HTMLElement) {
		k.Meaning = strings.TrimSpace(element.Text)
	})

	//Kunyomi
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-10.columns > div > div.small-12.large-7.columns.kanji-details__main > div.kanji-details__main-readings > dl.dictionary_entry.kun_yomi > dd", func(element *colly.HTMLElement) {
		element.ForEach("a", func(i int, element *colly.HTMLElement) {
			k.Kunyomi = append(k.Kunyomi, element.Text)
		})
	})

	//Kunyomi Examples
	c.OnHTML("#result_area > div > div:nth-child(3) > div.small-12.large-10.columns > div > div > div.row.compounds > div:nth-child(2) > ul", func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			k.KunyomiExamples = append(k.KunyomiExamples, parseExample(element.Text))
		})
	})

	//Onyomi
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-10.columns > div > div.small-12.large-7.columns.kanji-details__main > div.kanji-details__main-readings > dl.dictionary_entry.on_yomi > dd", func(element *colly.HTMLElement) {
		element.ForEach("a", func(i int, element *colly.HTMLElement) {
			k.Onyomi = append(k.Onyomi, element.Text)
		})
	})

	//Onyomi Examples
	c.OnHTML("#result_area > div > div:nth-child(3) > div.small-12.large-10.columns > div > div > div.row.compounds > div:nth-child(1) > ul", func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			k.OnyomiExamples = append(k.OnyomiExamples, parseExample(element.Text))
		})
	})

	//Radical
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-2.columns > div > div:nth-child(2) > div:nth-child(2) > dl > dd > span", func(element *colly.HTMLElement) {
		var r Radical
		r.Meaning = strings.TrimSpace(element.ChildText("span"))
		r.Symbol = strings.TrimSpace(strings.Trim(strings.TrimSpace(element.Text), r.Meaning))
		k.Radical = &r
	})

	//Parts
	c.OnHTML("#result_area > div > div:nth-child(1) > div.small-12.large-2.columns > div > div:nth-child(2) > div:nth-child(3) > dl > dd", func(element *colly.HTMLElement) {
		element.ForEach("a", func(i int, element *colly.HTMLElement) {
			k.Parts = append(k.Parts, element.Text)
		})
	})

	//Stroke Order Diagram
	k.StrokeOrderDiagram = fmt.Sprintf("%v%v_frames.png", j.StrokeOrderDiagramBaseUri, []rune(kanji)[0])

	//Jisho Page
	k.JishoUri = u

	c.Visit(u)
	if !k.found {
		return nil, &KanjiNotFound{Arg: kanji}
	}
	return &k, nil
}
