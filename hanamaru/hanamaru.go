package hanamaru

import (
	"github.com/bwmarrin/discordgo"
	"hanamaru/hanamaru/voice"
	"log"
	"strings"
)

type Hanamaru struct {
	prefix string
	*discordgo.Session
	VoiceContext *voice.Context
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

	voiceContext := voice.NewContext()

	log.Printf("Ready and logged in as %v zura!", s.State.User.Username+"#"+s.State.User.Discriminator)

	return &Hanamaru{prefix, s, voiceContext}
}

func (h *Hanamaru) AddCommand(cmd *Command) {
	var handleFunc = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if !strings.HasPrefix(m.Content, h.prefix+cmd.Name) {
			return
		}
		if m.Author.ID == s.State.User.ID || m.Author.Bot {
			return
		}
		argsString := strings.TrimPrefix(m.Content, h.prefix+cmd.Name)
		args := ParseArgs(argsString)
		ctx := &Context{
			Session:       s,
			MessageCreate: m,
			Args:          args,
			VoiceContext:  h.VoiceContext,
		}
		err := cmd.Exec(ctx)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "ERROR: "+err.Error())
		}
	}
	h.AddHandler(handleFunc)
}

func (h *Hanamaru) Close() {
	for _, vc := range h.VoiceContext.VCs {
		vc.Disconnect()
	}
	err := h.Session.Close()
	if err != nil {
		log.Fatalf("Failed to exit the bot correctly: %v", err)
	}
}
