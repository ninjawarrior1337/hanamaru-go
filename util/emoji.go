package util

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
)

var emojiRegex = regexp.MustCompile("<(a)?:([^><]+):(\\d+)>")

func ParseEmojis(input string) []discordgo.Emoji {
	var emoji = make([]discordgo.Emoji, 0)
	matches := emojiRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		emoji = append(emoji, discordgo.Emoji{Animated: match[1] != "", Name: match[2], ID: match[3]})
	}
	return emoji
}
