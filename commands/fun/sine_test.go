package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeSpaces(t *testing.T) {
	var integral = 0
	for i := 0; i <= SpacesMax*2; i++ {
		integral += len(computeSpaces(i, SpacesMax))
		t.Log(len(computeSpaces(i, SpacesMax)))
	}
	assert.LessOrEqual(t, integral, 2500)
}
