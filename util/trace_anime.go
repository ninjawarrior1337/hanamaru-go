package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"net/http"
)

const TRACE_URL = "https://api.trace.moe/search"

func TraceAnime(i image.Image) (*TraceMoeResp, error) {
	body := &bytes.Buffer{}

	mp := multipart.NewWriter(body)
	part, err := mp.CreateFormFile("image", "blob")
	if err != nil {
		return nil, fmt.Errorf("failed to create form part: %v", err)
	}
	jpeg.Encode(part, i, &jpeg.Options{Quality: 80})
	mp.Close()

	resp, err := http.Post(TRACE_URL, mp.FormDataContentType(), body)
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
