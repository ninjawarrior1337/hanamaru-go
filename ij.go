// +build ij

package main

import (
	"hanamaru/commands/image"
	"hanamaru/events"
)

func init() {
	commands = append(commands, image.Bishop)
	optionalEvents = append(optionalEvents, events.RepeatMessage)
}