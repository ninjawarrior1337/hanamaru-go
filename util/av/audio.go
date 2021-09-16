package av

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
)

func AddAudioToImage(img image.Image, audioFileName string) *bytes.Buffer {
	outBuf := bytes.NewBuffer([]byte{})
	audioF, close := createTempFile(audioFileName, Audio)
	defer close()

	ffmpegCmd, stdin, stdout := ffmpegCommand("-i", "pipe:0", "-i", audioF.Name(), "-c:v", "libvpx-vp9", "-c:a", "copy", "-f", "webm", "pipe:1")

	ffmpegCmd.Start()

	func() {
		defer stdin.Close()
		jpeg.Encode(stdin, img, nil)
	}()
	defer stdout.Close()
	io.Copy(outBuf, stdout)

	ffmpegCmd.Wait()

	return outBuf
}
