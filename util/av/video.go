package av

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
)

func OverlayImage(img image.Image, videoFileName string, startF, endF int) *bytes.Buffer {
	outBuf := &bytes.Buffer{}
	videoF, c := createTempFile(videoFileName, Video)
	defer c()

	ffmpegCmd, stdin, stdout := ffmpegCommand(
		"-i", "pipe:0",
		"-i", videoF.Name(),
		"-c:v", "libvpx-vp9",
		"-filter_complex", fmt.Sprintf("[1:v][0:v]overlay=enable=between(n\\, %v\\, %v)[out]", startF, endF),
		"-threads", "8",
		"-cpu-used", "4",
		"-tile-columns", "6",
		"-map", "[out]",
		"-map", "1:a",
		"-c:a", "copy",
		"-f", "webm",
		"-report",
		"pipe:1")

	ffmpegCmd.Start()

	func() {
		defer stdin.Close()
		jpeg.Encode(stdin, img, nil)
	}()
	defer stdout.Close()
	io.Copy(outBuf, stdout)

	return outBuf
}
