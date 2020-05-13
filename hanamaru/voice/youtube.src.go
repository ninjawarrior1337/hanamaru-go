package voice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/rylio/ytdl"
)

type YoutubeSrc struct {
	YtUrl string
	ec    *dca.EncodeSession
}

func (s *YoutubeSrc) Play(vc *discordgo.VoiceConnection) (*dca.StreamingSession, error) {
	videoInfo, err := ytdl.GetVideoInfo(s.YtUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get video info: %v", err)
	}

	format := videoInfo.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]
	downloadURL, err := videoInfo.GetDownloadURL(format)
	if err != nil {
		return nil, fmt.Errorf("failed to get download url: %v", err)
	}

	s.ec, err = dca.EncodeFile(downloadURL.String(), defaultOptions)
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
