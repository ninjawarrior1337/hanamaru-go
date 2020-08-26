package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
	"regexp"
)

var emojiRegex = regexp.MustCompile("")

var BigEmoji = &framework.EventListener{
	Name: "BigEmoji",
	HandlerConstructor: func(h *framework.Hanamaru) interface{} {
		return func(s *discordgo.Session, m *discordgo.MessageCreate) {
			if m.Author.Bot {
				return
			}
			emojis := util.ParseEmojis(m.Content)
			if len(emojis) == 1 && isSingleEmoji(m.Content) {
				s.ChannelMessageSendEmbed(m.ChannelID, constructEmojiEmbed(emojis[0], m.Author))
				s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
			}
		}
	},
}

func isSingleEmoji(input string) bool {
	matched, _ := regexp.MatchString("^<(a)?:([^><]+):(\\d+)>$", input)
	return matched
}

func constructEmojiEmbed(emoji discordgo.Emoji, user *discordgo.User) *discordgo.MessageEmbed {
	emb := &discordgo.MessageEmbed{
		URL:         "",
		Type:        "",
		Title:       "",
		Description: "",
		Timestamp:   "",
		Color:       0x3399ff,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    user.Username + "#" + user.Discriminator,
			IconURL: user.AvatarURL("256"),
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "",
		},
		Thumbnail: nil,
		Video:     nil,
		Provider:  nil,
		Author:    nil,
		Fields:    nil,
	}
	if emoji.Animated {
		emb.Image.URL = fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.gif?v=1", emoji.ID)
	} else {
		emb.Image.URL = fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.png?v=1", emoji.ID)
	}
	return emb
}
