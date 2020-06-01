package voice

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

func NewContext() *Context {
	ytdl, _ := NewYTDL("")
	return &Context{
		Queues: make(map[string]*Queue),
		VCs:    make(map[string]*discordgo.VoiceConnection),
		Ytdl:   ytdl,
	}
}

type Context struct {
	Queues map[string]*Queue
	VCs    map[string]*discordgo.VoiceConnection
	Ytdl   *Ytdl
}

type Playable interface {
	Play(vc *discordgo.VoiceConnection) (*dca.StreamingSession, error)
}
