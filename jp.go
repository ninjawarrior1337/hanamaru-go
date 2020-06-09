// +build jp

package main

import "github.com/ninjawarrior1337/hanamaru-go/commands/jp"

func init() {
	commands = append(commands, jp.Romaji, jp.Pitch)
}
