package voice

import (
	"encoding/json"
	"errors"
	"os/exec"
	"strings"
)

type Ytdl struct {
	Path string
}

func NewYTDL(path string) (*Ytdl, error) {
	if path == "" {
		if path = findYtdl(); path == "" {
			return nil, errors.New("cannot find youtubedl automatically, please specify a path")
		}
	}
	return &Ytdl{Path: path}, nil
}

func findYtdl() string {
	path, err := exec.LookPath("youtube-dl")
	if err != nil {
		return ""
	} else {
		return path
	}
}

func (y *Ytdl) GetVideoInfo(video string) (*Video, error) {
	var cmd *exec.Cmd
	if strings.HasPrefix(video, "http") {
		cmd = exec.Command("youtube-dl", "-j", video)
	} else {
		cmd = exec.Command("youtube-dl", "-j", "https://youtube.com/watch?v="+video)
	}
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	var v Video
	json.NewDecoder(stdOut).Decode(&v)
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	return &v, nil
}
