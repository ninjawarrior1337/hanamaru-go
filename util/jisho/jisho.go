package jisho

// This might spin off into its own package at some point.

type Jisho struct {
	BaseScrapeUri             string
	BaseApiUri                string
	StrokeOrderDiagramBaseUri string
}

func NewJisho() *Jisho {
	baseS := "https://jisho.org/search"
	return &Jisho{
		BaseScrapeUri:             baseS,
		BaseApiUri:                "https://jisho.org/api/v1/search/words?keyword=",
		StrokeOrderDiagramBaseUri: "http://classic.jisho.org/static/images/stroke_diagrams/",
	}
}
