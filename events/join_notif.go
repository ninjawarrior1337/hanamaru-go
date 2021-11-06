//go:build ij

package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var JoinNotif = &framework.EventListener{
	Name: "JoinNotif",
	HandlerConstructor: func(h *framework.Hanamaru) interface{} {
		return func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
			if m.GuildID != "407230022743621642" {
				return
			}
			msg := "https://en.wikipedia.org/wiki/Sakoku"
			guild, _ := s.Guild(m.GuildID)
			s.ChannelMessageSend(guild.SystemChannelID, msg)
		}
	},
}
