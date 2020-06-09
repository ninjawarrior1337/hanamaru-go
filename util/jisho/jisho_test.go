package jisho

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestJisho_SearchForKanji(t *testing.T) {
	j := NewJisho()
	spew.Dump(j.SearchForKanji("君"))
	spew.Dump(j.SearchForKanji("謎"))
}
