package events

import "github.com/bwmarrin/discordgo"

type Sent struct {
	Content string
	User    *discordgo.User
}

var lastTwoMessages = make(map[string][]Sent)

var RepeatMessage = func(s *discordgo.Session, m *discordgo.MessageCreate) {
	sent := Sent{Content: m.Content, User: m.Author}
	if len(lastTwoMessages[m.ChannelID]) < 2 {
		lastTwoMessages[m.ChannelID] = append(lastTwoMessages[m.ChannelID], Sent{})
		lastTwoMessages[m.ChannelID] = append(lastTwoMessages[m.ChannelID], sent)
	}
	lastTwoMessages[m.ChannelID][0], lastTwoMessages[m.ChannelID][1] = lastTwoMessages[m.ChannelID][1], sent

	if lastTwoMessages[m.ChannelID][0].User.Bot || lastTwoMessages[m.ChannelID][1].User.Bot {
		return
	}

	if lastTwoMessages[m.ChannelID][0].User.ID == lastTwoMessages[m.ChannelID][1].User.ID {
		return
	}

	if lastTwoMessages[m.ChannelID][0].Content == lastTwoMessages[m.ChannelID][1].Content {
		s.ChannelMessageSend(m.ChannelID, lastTwoMessages[m.ChannelID][0].Content)
	}
}
