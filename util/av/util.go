package av

import (
	"embed"
	"io"
	"os"
	"os/exec"
)

//go:embed assets/*
var avAssets embed.FS

type MediaType string

const (
	Audio MediaType = "audio"
	Video MediaType = "video"
)

func ffmpegCommand(args ...string) (*exec.Cmd, io.WriteCloser, io.ReadCloser) {
	ffmpegCmd := exec.Command("ffmpeg", args...)
	stdin, _ := ffmpegCmd.StdinPipe()
	stdout, _ := ffmpegCmd.StdoutPipe()

	return ffmpegCmd, stdin, stdout
}

func createTempFile(name string, mt MediaType) (file *os.File, close func()) {
	fileExt := ""
	switch mt {
	case Audio:
		fileExt = "ogg"
	case Video:
		fileExt = "webm"
	}
	embededFile, _ := avAssets.Open("assets/" + name)
	defer embededFile.Close()
	f, _ := os.CreateTemp("", "hanamaru_tmp_file.*."+fileExt)
	io.Copy(f, embededFile)
	return f, func() {
		f.Close()
		os.Remove(f.Name())
	}
}
