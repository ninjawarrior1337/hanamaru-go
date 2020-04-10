package hanamaru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArgs(t *testing.T) {
	// Test 1
	argStr := `fortnite a`
	arr := ParseArgs(argStr)
	assert.Equal(t, []string{"fortnite", "a"}, arr)

	// Test 2
	argStr = `"fort nite" a`
	arr = ParseArgs(argStr)
	assert.Equal(t, []string{"fort nite", "a"}, arr)

	// Test 3
	argStr = "https://www.youtube.com/watch?v=7WnzcgpaZ2s"
	arr = ParseArgs(argStr)
	assert.Equal(t, []string{"https://www.youtube.com/watch?v=7WnzcgpaZ2s"}, arr)

	argStr = "1 1"
	arr = ParseArgs(argStr)
	assert.Equal(t, []string{"1", "1"}, arr)

	argStr = "11"
	arr = ParseArgs(argStr)
	assert.Equal(t, []string{"11"}, arr)
}
