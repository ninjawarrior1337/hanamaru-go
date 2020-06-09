// +build ij

package main

import (
	"github.com/ninjawarrior1337/hanamaru-go/commands/image"
	"github.com/ninjawarrior1337/hanamaru-go/events"
)

func init() {
	commands = append(commands, image.Bishop)
	optionalEvents = append(optionalEvents, events.RepeatMessage)
}
