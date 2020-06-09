package jisho

// This might spin off into its own package at some point.

type Jisho struct {
	BaseScrapeURI             string
	StrokeOrderDiagramBaseUri string
}

func NewJisho() *Jisho {
	baseS := "https://jisho.org/search"
	return &Jisho{
		BaseScrapeURI:             baseS,
		StrokeOrderDiagramBaseUri: "http://classic.jisho.org/static/images/stroke_diagrams/",
	}
}
