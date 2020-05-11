package util

import (
	"os"
	"testing"
)

func PerformOnlyCI(t *testing.T) {
	if _, isCI := os.LookupEnv("CI"); isCI {
		t.SkipNow()
	}
}
