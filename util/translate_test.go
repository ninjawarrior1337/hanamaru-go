package util

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTranslate(t *testing.T) {
	tr, err := Translate("hello", "en", "ja")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, tr, "こんにちは")
}
