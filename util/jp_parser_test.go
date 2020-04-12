// +build jp

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseJapanese(t *testing.T) {
	tests := []string{
		"日本語でしゃべません",
		"タイラー",
	}
	expected := []string{
		"nihongo de shabe mase n",
		"taira-",
	}
	for i, test := range tests {
		actual := ParseJapanese(test)
		assert.Equal(t, expected[i], actual)
	}
}
