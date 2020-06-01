package util

import (
	"fmt"
	"testing"
)

func TestParseNhentai(t *testing.T) {
	PerformNotCI(t)
	n, _ := ParseNhentai(308389)

	fmt.Println(n)
	fmt.Println(n.Tags)
}
