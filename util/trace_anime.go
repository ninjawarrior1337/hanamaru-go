package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
)

const TRACE_URL = "https://trace.moe/api/search"

func imageToJpgBase64(i image.Image) (string, error) {
	imgBuf := &bytes.Buffer{}
	b64Buf := &bytes.Buffer{}
	if err := jpeg.Encode(imgBuf, i, &jpeg.Options{Quality: 90}); err != nil {
		return "", fmt.Errorf("failed to encode image in jpg: %v", err)
	}
	if _, err := base64.NewEncoder(base64.RawStdEncoding, b64Buf).Write(imgBuf.Bytes()); err != nil {
		return "", fmt.Errorf("failed to encode image in base64: %v", err)
	}
	return b64Buf.String(), nil
}

func TraceAnime(i image.Image) (*TraceMoeResp, error) {
	body := &bytes.Buffer{}
	b64, err := imageToJpgBase64(i)
	if err != nil {
		return nil, fmt.Errorf("failed to preprocess image: %v", err)
	}
	body.Write([]byte(fmt.Sprintf(`{"image": "%v"}`, b64)))
	resp, err := http.Post(TRACE_URL, "application/json", body)
	if err != nil {
		return nil, fmt.Errorf("request to trace.moe failed: %v", err)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 429 {
			return nil, fmt.Errorf("you have been rate limited by trace.moe")
		}
		return nil, fmt.Errorf("trace.moe failed to decode this image: resp code %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	var t TraceMoeResp
	json.NewDecoder(resp.Body).Decode(&t)

	return &t, nil
}
