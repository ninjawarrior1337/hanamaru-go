package hanamaru

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type Hanamaru struct {
	prefix string
	*discordgo.Session
}

func New(t, prefix string) (bot *Hanamaru) {
	s, err := discordgo.New(t)

	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	err = s.Open()
	if err != nil {
		log.Fatalf("Failed to open webhook: %v", err)
	}

	return &Hanamaru{prefix, s}
}

func (h *Hanamaru) AddCommand(cmd *Command) {
	var handleFunc = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if !strings.HasPrefix(m.Content, h.prefix+cmd.Name) {
			return
		}
		argsString := strings.TrimPrefix(m.Content, h.prefix+cmd.Name)
		args := ParseArgs(argsString)
		ctx := &Context{
			Session:       s,
			MessageCreate: m,
			Args:          args,
		}
		err := cmd.Exec(ctx)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "ERROR: "+err.Error())
		}
	}
	h.AddHandler(handleFunc)
}

func (h *Hanamaru) Close() {
	err := h.Session.Close()
	if err != nil {
		log.Fatalf("Failed to exit the bot correctly: %v", err)
	}
}
