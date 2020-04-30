package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var Boomer = func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	fmt.Print(r.Emoji)
	if r.Emoji.ID == "" {
		fmt.Print(r.Emoji)
		s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.ID, r.UserID)

		s.MessageReactionAdd(r.ChannelID, r.MessageID, "🆗")
	}
}
