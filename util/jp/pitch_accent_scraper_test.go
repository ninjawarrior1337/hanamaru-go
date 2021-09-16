//go:build jp

package jp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapePitchAccent(t *testing.T) {
	if _, isCI := os.LookupEnv("CI"); isCI {
		return
	}
	tests := []string{
		"おはようございます",
		"意味わかない",
	}
	expectedPitch := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1},
	}
	expectedPhrase := []string{
		"おはようございます",
		"いみわかない",
	}
	for i, test := range tests {
		actualPhrase, actualPitch := ScrapePitchAccent(test)

		assert.Equal(t, expectedPhrase[i], actualPhrase)
		assert.Equal(t, expectedPitch[i], actualPitch)
	}
}
