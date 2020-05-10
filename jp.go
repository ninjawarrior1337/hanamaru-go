// +build jp

package main

import "hanamaru/commands/jp"

func init() {
	commands = append(commands, jp.Romaji, jp.Pitch)
}
