package voice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

var defaultOptions *dca.EncodeOptions

func init() {
	defaultOptions = dca.StdEncodeOptions
	defaultOptions.RawOutput = true
	defaultOptions.Bitrate = 96
	defaultOptions.Application = "lowdelay"
}

type StaticFile struct {
	FilePath string
	ec       *dca.EncodeSession
}

func (s *StaticFile) Play(vc *discordgo.VoiceConnection) (*dca.StreamingSession, error) {
	var err error
	s.ec, err = dca.EncodeFile(s.FilePath, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to play file: %v: %v", s.FilePath, err)
	}
	fmt.Println(s.ec.Stats())
	done := make(chan error)
	stream := dca.NewStream(s.ec, vc, done)
	err = <-done
	fmt.Println(err)
	fmt.Println(s.ec.FFMPEGMessages())
	return stream, nil
}
