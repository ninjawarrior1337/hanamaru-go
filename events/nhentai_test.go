package events

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStringWithSixDigits(t *testing.T) {
	testString := "fortnite456693 456692 456691"

	numbers, err := ParseStringWithSixDigits(testString)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, []int{456692, 456691}, numbers)

	testString = "https://cdn.discordapp.com/attachments/407230022743621644/698664289279672480/image0.png"
	numbers, err = ParseStringWithSixDigits(testString)

	assert.Equal(t, err, fmt.Errorf("no matches"))

}
