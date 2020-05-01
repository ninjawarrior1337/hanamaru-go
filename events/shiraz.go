// +build ij

package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"hanamaru/util"
)

var shiraz = []string{"🆗"}

func init() {
	shiraz = append(shiraz, util.MustMapToEmoji("shiraz")...)
}

var Shiraz = func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.Emoji.Name == "🇮🇳" {
		fmt.Print(r.Emoji)
		s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.APIName(), r.UserID)

		for _, emojiId := range shiraz {
			s.MessageReactionAdd(r.ChannelID, r.MessageID, emojiId)
		}
	}
}
