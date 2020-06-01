package voice

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
	"github.com/jonas747/dca"
	"strings"
)

type YoutubeSrc struct {
	YtUrl string
	Ytdl  *Ytdl
	ec    *dca.EncodeSession
}

func NewYTSrc(url string, ytdl *Ytdl) *YoutubeSrc {
	return &YoutubeSrc{YtUrl: url, Ytdl: ytdl}
}

func (s *YoutubeSrc) Play(vc *discordgo.VoiceConnection) (*dca.StreamingSession, error) {
	videoInfo, err := s.Ytdl.GetVideoInfo(s.YtUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get video info: %v", err)
	}
	var format Formats
	spew.Dump(videoInfo.Formats)
	for _, f := range videoInfo.Formats {
		if strings.ContainsAny(f.Format, "audio only") && f.Abr > 50 {
			format = f
			break
		}
	}
	if format.URL == "" {
		return nil, errors.New("failed to get url")
	}
	s.ec, err = dca.EncodeFile(format.URL, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create encoder: %v", err)
	}

	done := make(chan error)
	stream := dca.NewStream(s.ec, vc, done)
	err = <-done
	fmt.Println(err)
	fmt.Println(s.ec.FFMPEGMessages())
	return stream, err
}
