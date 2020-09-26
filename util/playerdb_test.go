package util

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestLookupMinecraft(t *testing.T) {
	SkipCI(t)
	p, err := LookupMinecraft("ninjawarrior1337")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, p.Id, "7ce98ec4-496c-4ab0-968b-280c750f423b")
}
