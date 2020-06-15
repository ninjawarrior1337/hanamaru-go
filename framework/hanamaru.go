package framework

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework/voice"
	bolt "go.etcd.io/bbolt"
	"log"
	"strings"
)

type Hanamaru struct {
	prefix  string
	ownerid string
	*discordgo.Session
	VoiceContext *voice.Context
	commands     []*Command
	Db           *bolt.DB
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

func HasPermission(s *discordgo.Session, member *discordgo.Member, permission int) (bool, error) {
	for _, rid := range member.Roles {
		role, err := s.State.Role(member.GuildID, rid)
		if err != nil {
			return false, err
		}
		if role.Permissions&permission != 0 {
			return true, nil
		}
	}
	return false, nil
}

type ErrAddCommand struct {
	cmd *Command
	error
}

func (e ErrAddCommand) Error() string {
	return fmt.Sprintf("failed to add command: %v: reason: %v", e.cmd.Name, e)
}

func (h *Hanamaru) AddCommand(cmd *Command) error {
	if cmd.Name == "" {
		panic("A command must not have an empty name!")
	}
	if cmd.Setup != nil {
		err := cmd.Setup()
		if err != nil {
			return ErrAddCommand{
				cmd:   cmd,
				error: err,
			}
		}
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
		if ok, _ := HasPermission(s, m.Member, cmd.PermissionRequired); cmd.PermissionRequired > 0 && !ok {
			s.ChannelMessageSend(m.ChannelID, "ERROR: You don't have the required permissions to run this command")
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
	return nil
}

func (h *Hanamaru) EnableHelpCommand() {
	h.AddCommand(help)
}

func (h *Hanamaru) EnableDB() {

}

func (h *Hanamaru) Close() {
	for _, vc := range h.VoiceContext.VCs {
		vc.Disconnect()
	}
	err := h.Session.Close()
	if err != nil {
		log.Fatalf("Failed to exit the bot correctly: %v", err)
	}
	if h.Db != nil {
		h.Db.Close()
	}
}
