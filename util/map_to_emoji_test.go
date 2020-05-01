package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapToEmoji(t *testing.T) {
	tests := []string{
		"abo",
		"aboabo",
		"aboaboo",
		"ethan",
	}

	expected := [][]string{
		{"🇦", "🇧", "🇴"},
		{"🇦", "🇧", "🇴", "🅰️", "🅱️", "🅾️"},
		{},
		{"🇪", "🇹", "🇭", "🇦", "🇳"},
	}

	for i, test := range tests {
		emoji, err := MapToEmoji(test)
		if i == 2 {
			assert.Equal(t, EmojiConversionError{"aboaboo", 6}, err)
			continue
		}
		assert.Equal(t, expected[i], emoji)
	}
}
