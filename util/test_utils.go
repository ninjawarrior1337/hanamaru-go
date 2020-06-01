package util

import (
	"os"
	"testing"
)

func PerformNotCI(t *testing.T) {
	if _, isCI := os.LookupEnv("CI"); isCI {
		t.SkipNow()
	}
}
