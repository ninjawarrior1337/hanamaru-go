package util

import (
	"fmt"
	"testing"
)

func TestParseNhentai(t *testing.T) {
	n, _ := ParseNhentai(308389)

	fmt.Println(n)
	fmt.Println(n.Tags)
}
