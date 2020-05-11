package hanamaru

import (
	"github.com/bwmarrin/discordgo"
	bolt "go.etcd.io/bbolt"
	"hanamaru/hanamaru/voice"
	"log"
	"strings"
)

type Hanamaru struct {
	prefix  string
	ownerid string
	*discordgo.Session
	VoiceContext *voice.Context
	commands     []*Command
	db           *bolt.DB
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

	return &Hanamaru{
		prefix:       prefix,
		Session:      s,
		VoiceContext: voiceContext,
		ownerid:      "",
		commands:     []*Command{},
	}
}

func (h *Hanamaru) SetOwner(id string) {
	h.ownerid = id
}

func (h *Hanamaru) AddCommand(cmd *Command) {
	if cmd.Name == "" {
		panic("A command must not have an empty name!")
	}
	var handleFunc = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if !strings.HasPrefix(m.Content, h.prefix+cmd.Name) {
			return
		}
		if m.Author.ID == s.State.User.ID || m.Author.Bot {
			return
		}
		if cmd.OwnerOnly && h.ownerid != m.Author.ID {
			s.ChannelMessageSend(m.ChannelID, "ERROR: You must be the owner of this instance to run this command")
			return
		}
		argsString := strings.TrimPrefix(m.Content, h.prefix+cmd.Name)
		args := ParseArgs(argsString)
		ctx := &Context{
			Session:       s,
			MessageCreate: m,
			Args:          args,
			VoiceContext:  h.VoiceContext,
			Hanamaru:      h,
		}
		err := cmd.Exec(ctx)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "ERROR: "+err.Error())
		}
	}
	h.AddHandler(handleFunc)
	h.commands = append(h.commands, cmd)
}

func (h *Hanamaru) EnableHelpCommand() {
	h.AddCommand(help)
}

func (h *Hanamaru) Close() {
	for _, vc := range h.VoiceContext.VCs {
		vc.Disconnect()
	}
	err := h.Session.Close()
	if err != nil {
		log.Fatalf("Failed to exit the bot correctly: %v", err)
	}
	if h.db != nil {
		h.db.Close()
	}
}
