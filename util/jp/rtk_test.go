//go:build jp

package jp

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewRtkDatabase(t *testing.T) {
	rtk := NewRtkDatabase()
	spew.Dump(rtk.Entries[1])
}
