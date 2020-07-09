package jisho

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ninjawarrior1337/hanamaru-go/util"
	"testing"
)

func TestJisho_SearchForKanji(t *testing.T) {
	util.PerformNotCI(t)
	j := NewJisho()
	spew.Dump(j.SearchForKanji("君"))
	spew.Dump(j.SearchForKanji("謎"))
	spew.Dump(j.SearchForKanji("宰"))
}
