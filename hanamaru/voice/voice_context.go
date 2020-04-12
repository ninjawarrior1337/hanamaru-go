package voice

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

func NewContext() *Context {
	return &Context{
		Queues: make(map[string]*Queue),
		VCs:    make(map[string]*discordgo.VoiceConnection),
	}
}

type Context struct {
	Queues map[string]*Queue
	VCs    map[string]*discordgo.VoiceConnection
}

type Playable interface {
	Play(vc *discordgo.VoiceConnection) (*dca.StreamingSession, error)
}
