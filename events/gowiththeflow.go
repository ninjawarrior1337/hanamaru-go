// +build ij

package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

const GoWithTheFlowLimit = 3

type Sent struct {
	Content string
	User    *discordgo.User
}

var lastTwoMessages = make(map[string][]Sent)

var RepeatMessage = &framework.EventListener{
	Name: "Repeat Message",
	HandlerConstructor: func(h *framework.Hanamaru) interface{} {
		return func(s *discordgo.Session, m *discordgo.MessageCreate) {
			sent := Sent{Content: m.Content, User: m.Author}
			sentArr := lastTwoMessages[m.ChannelID]
			shiftSent(&sentArr, sent, GoWithTheFlowLimit)

			if didGoWithTheFlow(sentArr) {
				s.ChannelMessageSend(m.ChannelID, lastTwoMessages[m.ChannelID][0].Content)
			}

			lastTwoMessages[m.ChannelID] = sentArr
		}
	},
}

func didGoWithTheFlow(arr []Sent) bool {
	testString := arr[0].Content
	for _, s := range arr {
		if s.User.Bot {
			return false
		}
		if s.Content != testString {
			return false
		}
	}
	if !areIdsUnique(arr) {
		return false
	}
	return true
}

func areIdsUnique(arr []Sent) bool {
	for outPtr, outerS := range arr {
		for inPtr, innerS := range arr {
			if outPtr == inPtr {
				continue
			}
			if outerS.User.ID == innerS.User.ID {
				return false
			}
		}
	}
	return true
}

func shiftSent(arr *[]Sent, msg Sent, limit int) {
	for i := len(*arr) - 1; i >= 0; i-- {
		s := (*arr)[i]
		if len(*arr) <= i+1 {
			*arr = append(*arr, s)
			continue
		}
		(*arr)[i+1] = s
	}
	if len(*arr) == 0 {
		*arr = append(*arr, msg)
	} else {
		(*arr)[0] = msg
	}

	if len(*arr) > limit {
		*arr = (*arr)[0:limit]
	}
}
