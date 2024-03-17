package events

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	hdb "github.com/ninjawarrior1337/hanamaru-go/db"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var (
	ErrInvalidAwardType = errors.New("invalid award type")
)

type AwardType string

func (a AwardType) IsValid() bool {
	if len(strings.Split(string(a), "_")) == 2 {
		if strings.Split(string(a), "_")[1] == "award" {
			return true
		}
	}
	return false
}

func (a AwardType) Name() (string, error) {
	if !a.IsValid() {
		return "", ErrInvalidAwardType
	}

	return a.MustName(), nil
}

func (a AwardType) MustName() string {
	return strings.ToLower(strings.Split(string(a), "_")[0])
}

var AwardsAddHandler = &framework.EventListener{
	Name: "Award Add",
	HandlerConstructor: func(h *framework.Hanamaru) interface{} {
		return func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
			award := AwardType(r.Emoji.Name)
			m, _ := s.ChannelMessage(r.ChannelID, r.MessageID)
			if name, err := award.Name(); err == nil {
				hdb.MutateAwardCount(h.Db, r.GuildID, m.Author.ID, name, hdb.AwardIncrement)
			}
		}
	},
}

var AwardsRemoveHandler = &framework.EventListener{
	Name: "Award Remove",
	HandlerConstructor: func(h *framework.Hanamaru) interface{} {
		return func(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
			award := AwardType(r.Emoji.Name)
			m, _ := s.ChannelMessage(r.ChannelID, r.MessageID)
			if name, err := award.Name(); err == nil {
				hdb.MutateAwardCount(h.Db, r.GuildID, m.Author.ID, name, hdb.AwardDecrement)
			}
		}
	},
}
