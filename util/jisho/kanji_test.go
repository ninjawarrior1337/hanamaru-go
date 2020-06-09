package jisho

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseExample(t *testing.T) {
	e := parseExample("後日 【ゴジツ】 in the future, another day, later")
	assert.Equal(t, &ReadingExample{
		"後日",
		"ゴジツ",
		"in the future, another day, later",
	}, e)
}
