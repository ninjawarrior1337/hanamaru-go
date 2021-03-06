package fun

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"strings"
)

var Nuke = &framework.Command{
	Name:               "nuke",
	PermissionRequired: 0,
	OwnerOnly:          true,
	Help:               "NEVER RUN THIS UNLESS U WANT TO NUKE THE SERVER",
	Exec: func(ctx *framework.Context) error {
		var outerr error
		ctx.Reply("Are you 100% sure you want to do this: (y or n)")
		ctx.Hanamaru.AddEventListenerOnce(&framework.EventListener{
			Name: "Nuke confirmer",
			HandlerConstructor: func(h *framework.Hanamaru) interface{} {
				return func(session *discordgo.Session, m *discordgo.MessageCreate) {
					if ctx.GuildID == m.GuildID && strings.Contains(m.Content, "y") && ctx.Author == m.Author {
						err := NukeGuild(ctx.Hanamaru.Session, ctx.GuildID)
						if err != nil {
							outerr = err
						}
					} else {
						ctx.Reply("Canceled")
					}
				}
			}})
		if outerr != nil {
			return outerr
		}
		return nil
	},
}

func NukeGuild(s *discordgo.Session, guildID string) error {
	g, err := s.Guild(guildID)
	if err != nil {
		return err
	}
	for _, channel := range g.Channels {
		s.ChannelDelete(channel.ID)
	}
	for _, user := range g.Members {
		if !user.User.Bot {
			s.GuildBanCreate(guildID, user.User.ID, 0)
		}
	}
	return nil
}
