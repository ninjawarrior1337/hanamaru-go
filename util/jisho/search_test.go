package jisho

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ninjawarrior1337/hanamaru-go/util"
	"testing"
)

var j *Jisho

func init() {
	j = NewJisho()
}

func TestJisho_SearchKeyword(t *testing.T) {
	util.SkipCI(t)
	spew.Dump(j.SearchKeyword("始まる"))
}
