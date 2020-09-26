package util

import (
	"os"
	"testing"
)

func SkipCI(t *testing.T) {
	if _, isCI := os.LookupEnv("CI"); isCI {
		t.SkipNow()
	}
}
