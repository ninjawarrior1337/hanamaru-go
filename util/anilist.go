package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetAnimeInfo(title string) (ALMedia, error) {
	body := ALRequest{Query: fmt.Sprintf(SearchQuery, title)}
	var b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	resp, err := http.Post("https://graphql.anilist.co", "application/json", b)
	if err != nil {
		return ALMedia{}, errors.New("failed to get info for requested show: " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return ALMedia{}, errors.New("failed to get info for requested show; Status code: " + string(resp.StatusCode))
	}
	var ar ALResponse
	json.NewDecoder(resp.Body).Decode(&ar)
	return ar.Data.Media, nil
}
