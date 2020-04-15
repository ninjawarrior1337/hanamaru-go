package util

import (
	"fmt"
	"os"
	"testing"
)

func TestParseNhentai(t *testing.T) {
	if _, isCI := os.LookupEnv("CI"); isCI {
		return
	}
	n, _ := ParseNhentai(308389)

	fmt.Println(n)
	fmt.Println(n.Tags)
}
