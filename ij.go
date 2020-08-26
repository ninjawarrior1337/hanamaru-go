// +build ij

package main

import (
	"github.com/ninjawarrior1337/hanamaru-go/commands/ij"
	"github.com/ninjawarrior1337/hanamaru-go/events"
)

func init() {
	commands = append(commands, ij.Bishop, ij.Schedule)
	optionalEvents = append(optionalEvents, events.RepeatMessage)
}
