// +build ij

package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"hanamaru/util"
)

var boomer = []string{"🆗"}

func init() {
	boomer = append(boomer, util.MustMapToEmoji("boomer")...)
}

var Boomer = func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.Emoji.Name == "💥" {
		fmt.Print(r.Emoji)
		s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.APIName(), r.UserID)

		for _, emojiId := range boomer {
			s.MessageReactionAdd(r.ChannelID, r.MessageID, emojiId)
		}
	}
}
