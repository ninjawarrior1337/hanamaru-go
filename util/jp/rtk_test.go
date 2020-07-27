//+build jp

package jp

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewRtkDatabase(t *testing.T) {
	rtk := NewRtkDatabase()
	spew.Dump(rtk.Entries[1])
}
