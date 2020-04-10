package voice

import (
	"github.com/bwmarrin/discordgo"
)

func NewContext() *Context {
	return &Context{
		Queues: map[string]*Queue{},
		VCs:    map[string]*discordgo.VoiceConnection{},
	}
}

type Context struct {
	Queues map[string]*Queue
	VCs    map[string]*discordgo.VoiceConnection
}

type Playable interface {
	Play(vc *discordgo.VoiceConnection) (chan error, error)
}
