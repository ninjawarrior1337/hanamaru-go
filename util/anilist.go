package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func GetAnimeInfoFromID(id int) (ALMedia, error) {
	req := ALRequest{Query: fmt.Sprintf(SearchQueryTemplate, "id", fmt.Sprintf(`%v`, id), "type", "ANIME")}
	return performRequest(req)
}

func GetAnimeInfoFromTitle(title string) (ALMedia, error) {
	req := ALRequest{Query: fmt.Sprintf(SearchQueryTemplate, "search", fmt.Sprintf(`"%v"`, title), "type", "ANIME")}
	return performRequest(req)
}

func GetMangaInfoFromTitle(title string) (ALMedia, error) {
	req := ALRequest{Query: fmt.Sprintf(SearchQueryTemplate, "search", fmt.Sprintf(`"%v"`, title), "type", "MANGA")}
	return performRequest(req)
}

func performRequest(req ALRequest) (ALMedia, error) {
	var b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(req)
	resp, err := http.Post("https://graphql.anilist.co", "application/json", b)
	if err != nil {
		return ALMedia{}, errors.New("failed to get info for requested show: " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return ALMedia{}, errors.New("failed to get info for requested show; status code: " + strconv.Itoa(resp.StatusCode))
	}
	var ar ALResponse
	json.NewDecoder(resp.Body).Decode(&ar)
	return ar.Data.Media, nil
}
