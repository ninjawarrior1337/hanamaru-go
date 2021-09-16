package framework

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework/voice"
	bolt "go.etcd.io/bbolt"
)

type Hanamaru struct {
	prefix  string
	ownerId string
	*discordgo.Session
	VoiceContext   *voice.Context
	commands       []*Command
	eventListeners []*EventListener
	Db             *bolt.DB
}

func New(token, prefix, ownerid string) (bot *Hanamaru) {
	s, err := discordgo.New(token)

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
		ownerId:      ownerid,
		commands:     []*Command{},
	}
}

func (h *Hanamaru) GetPrefix() string {
	return h.prefix
}

func (h *Hanamaru) GetOwnerID() string {
	return h.ownerId
}

func HasPermission(s *discordgo.Session, userID string, channelID string, reqPerm int64) (bool, error) {
	userPerms, err := s.State.UserChannelPermissions(userID, channelID)
	if err != nil {
		return false, err
	}

	if (userPerms | reqPerm) == userPerms {
		return true, nil
	}

	return false, nil
}

type ErrAddCommand struct {
	cmd    *Command
	reason string
	error
}

func (e ErrAddCommand) Error() string {
	return fmt.Sprintf("failed to add command: %v: reason: %v", e.cmd.Name, e.reason)
}

func (h *Hanamaru) AddCommand(cmd *Command) error {
	if cmd.Name == "" {
		panic("A command must not have an empty name!")
	}
	if cmd.Setup != nil {
		err := cmd.Setup()
		if err != nil {
			return ErrAddCommand{
				cmd:    cmd,
				reason: err.Error(),
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
		if cmd.OwnerOnly && h.ownerId != m.Author.ID {
			s.ChannelMessageSend(m.ChannelID, "ERROR: You must be the owner of this instance to run this command")
			return
		}
		if ok, _ := HasPermission(s, m.Author.ID, m.ChannelID, cmd.PermissionRequired); cmd.PermissionRequired > 0 && !ok {
			s.ChannelMessageSend(m.ChannelID, "ERROR: You don't have the required permissions to run this command")
			return
		}
		ctx := NewContext(h, cmd, m)

		h.Session.ChannelTyping(ctx.ChannelID)
		err := cmd.Exec(ctx)

		if err != nil {
			s.ChannelMessageSendReply(m.ChannelID, "ERROR: "+err.Error(), ctx.Reference())
		}
	}
	h.Session.AddHandler(handleFunc)
	h.commands = append(h.commands, cmd)
	return nil
}

// AddEventListener adds and event listener to the bot.
func (h *Hanamaru) AddEventListener(listener *EventListener) error {
	h.Session.AddHandler(listener.HandlerConstructor(h))
	h.eventListeners = append(h.eventListeners, listener)
	return nil
}

func (h *Hanamaru) AddEventListenerOnce(listener *EventListener) error {
	h.Session.AddHandlerOnce(listener.HandlerConstructor(h))
	return nil
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
	if h.Db != nil {
		h.Db.Close()
	}
}
