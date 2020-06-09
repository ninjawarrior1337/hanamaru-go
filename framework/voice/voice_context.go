package voice

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

func NewContext() *Context {
	ytdl, _ := NewYTDL()
	return &Context{
		Queues:        make(map[string]*Queue),
		QueueChannels: make(map[string]chan Playable),
		VCs:           make(map[string]*discordgo.VoiceConnection),
		Ytdl:          ytdl,
	}
}

type Context struct {
	Queues        map[string]*Queue
	QueueChannels map[string]chan Playable
	VCs           map[string]*discordgo.VoiceConnection
	Ytdl          *Ytdl
}

func (c *Context) JoinChannel(s *discordgo.Session, guildID, vcId, notifTextChannelID string) error {
	vc, err := s.ChannelVoiceJoin(guildID, vcId, false, false)
	if err != nil {
		return err
	}
	c.VCs[guildID] = vc

	var qChan = make(chan Playable, 1024)
	c.QueueChannels[guildID] = qChan
	go func() {
		defer delete(c.VCs, guildID)
		defer delete(c.QueueChannels, guildID)
		for p := range qChan {
			_, done, err := p.Play(vc)
			if err != nil {
				//TODO: Implement a way to get the title, probably via a method in the Playable interface.
				s.ChannelMessageSend(notifTextChannelID, fmt.Sprintf("Error playing: %v, Skipping", "this is supposed to be the title"))
				continue
			}
			select {
			case <-done:
			}
			p.Cleanup()
		}
	}()
	for {
		if vc.Ready {
			s.ChannelMessageSend(notifTextChannelID, "Joined!")
			break
		}
	}
	return nil
}

func (c *Context) LeaveChannel(guildID, vcId string) error {
	if vc, ok := c.VCs[guildID]; ok {
		if vc.ChannelID == vcId {
			vc.Disconnect()
			close(c.QueueChannels[guildID])
			return nil
		} else {
			return errors.New("i am not in this vc")
		}
	}
	return errors.New("cannot leave voice channel when not in one")
}

type Playable interface {
	Play(vc *discordgo.VoiceConnection) (*dca.StreamingSession, chan error, error)
	Cleanup()
}
