package jisho

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type SearchResp struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		Slug     string   `json:"slug"`
		IsCommon bool     `json:"is_common"`
		Tags     []string `json:"tags"`
		Jlpt     []string `json:"jlpt"`
		Japanese []struct {
			Word    string `json:"word"`
			Reading string `json:"reading"`
		} `json:"japanese"`
		Senses []struct {
			EnglishDefinitions []string `json:"english_definitions"`
			PartsOfSpeech      []string `json:"parts_of_speech"`
			Links              []string `json:"links"`
			Tags               []string `json:"tags"`
			Restrictions       []string `json:"restrictions"`
			SeeAlso            []string `json:"see_also"`
			Antonyms           []string `json:"antonyms"`
			Source             []string `json:"source"`
			Info               []string `json:"info"`
		} `json:"senses"`
		Attribution struct {
			Jmdict   bool `json:"jmdict"`
			Jmnedict bool `json:"jmnedict"`
			Dbpedia  bool `json:"dbpedia"`
		} `json:"attribution"`
	} `json:"data"`
}

func (j *Jisho) SearchKeyword(kw string) (*SearchResp, error) {
	resp, err := http.Get(j.BaseApiUri + url.QueryEscape(kw))
	if err != nil {
		return nil, errors.New("failed to reach jisho servers")
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("invalid request to jisho: " + fmt.Sprintf("%v", resp.StatusCode))
	}
	defer resp.Body.Close()
	var sr SearchResp
	json.NewDecoder(resp.Body).Decode(&sr)
	return &sr, nil
}
