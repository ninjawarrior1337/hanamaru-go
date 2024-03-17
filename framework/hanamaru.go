package framework

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework/voice"
)

type Hanamaru struct {
	Prefix  string
	ownerId string
	*discordgo.Session
	VoiceContext   *voice.Context
	Commands       []*Command
	EventListeners []*EventListener
	Db             *sql.DB
}

func New(token, prefix, ownerid string) (bot *Hanamaru) {
	s, err := discordgo.New(token)

	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsMessageContent

	err = s.Open()
	if err != nil {
		log.Fatalf("Failed to open webhook: %v", err)
	}

	voiceContext := voice.NewContext()

	log.Printf("Ready and logged in as %v zura!", s.State.User.Username+"#"+s.State.User.Discriminator)

	return &Hanamaru{
		Prefix:       prefix,
		Session:      s,
		VoiceContext: voiceContext,
		ownerId:      ownerid,
		Commands:     []*Command{},
	}
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
	reason error
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
				reason: err,
			}
		}
	}
	var handleFunc = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if !strings.HasPrefix(m.Content, h.Prefix+cmd.Name) {
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
	h.Commands = append(h.Commands, cmd)
	return nil
}

// AddEventListener adds and event listener to the bot.
func (h *Hanamaru) AddEventListener(listener *EventListener) error {
	h.Session.AddHandler(listener.HandlerConstructor(h))
	h.EventListeners = append(h.EventListeners, listener)
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
