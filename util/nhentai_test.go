package util

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestParseNhentai(t *testing.T) {
	SkipCI(t)
	//I randomly picked these numbers
	n, _ := ParseNhentai(308389)

	spew.Dump(n)
}
