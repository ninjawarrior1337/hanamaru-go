package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/util"
)

var expansionMap = map[string][]string{}

func init() {
	expansionMap["💥"] = append([]string{"🆗"}, util.MustMapToEmoji("boomer")...)
}

var ReactionExpansion = func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if expansion, ok := expansionMap[r.Emoji.Name]; ok {
		s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.APIName(), r.UserID)

		for _, emojiId := range expansion {
			s.MessageReactionAdd(r.ChannelID, r.MessageID, emojiId)
		}
	}
}
