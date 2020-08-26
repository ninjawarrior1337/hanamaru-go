package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShiftSent(t *testing.T) {
	arr := []Sent{}
	shiftSent(&arr, Sent{Content: "1"}, 3)
	assert.Equal(t, arr, []Sent{{Content: "1"}})

	shiftSent(&arr, Sent{Content: "2"}, 3)
	assert.Equal(t, arr, []Sent{{Content: "2"}, {Content: "1"}})

	shiftSent(&arr, Sent{Content: "3"}, 3)
	assert.Equal(t, arr, []Sent{{Content: "3"}, {Content: "2"}, {Content: "1"}})

	for _ = range [100]int{} {
		shiftSent(&arr, Sent{Content: "100"}, 3)
	}
	assert.LessOrEqual(t, cap(arr), 100)
}

func TestAreIdsUnique(t *testing.T) {
	arr := []Sent{
		{User: &discordgo.User{ID: "21"}},
		{User: &discordgo.User{ID: "22"}},
		{User: &discordgo.User{ID: "1"}},
	}
	assert.Equal(t, true, areIdsUnique(arr))
	arr = []Sent{
		{User: &discordgo.User{ID: "21"}},
		{User: &discordgo.User{ID: "100"}},
		{User: &discordgo.User{ID: "21"}},
	}
	assert.Equal(t, false, areIdsUnique(arr))
}
