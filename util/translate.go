package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"unicode/utf8"
)

type TranslateRequest struct {
	SourceText string
	SourceLang string
	TargetLang string
}

type TranslateResponse []interface{}

func Translate(st, sl, tl string) (string, error) {
	if sl == "" {
		sl = "auto"
	}
	if st == "" || tl == "" {
		return "", errors.New("all params must be non-empty strings")
	}
	if (sl != "auto" && utf8.RuneCountInString(sl) != 2) || utf8.RuneCountInString(tl) != 2 {
		return "", errors.New("source and target language must be the two letter iso code")
	}
	resp, err := http.Get(fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=%v&tl=%v&dt=t&q=%v", sl, tl, url.QueryEscape(st)))
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("failed to translate using google's translate api")
	}
	defer resp.Body.Close()
	var tr TranslateResponse
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		return "", errors.New("failed to decode translation response")
	}
	//I've been told not to do this, but it works sooooo..............if there is a better way to do this
	//please submit a pull req.
	return tr[0].([]interface{})[0].([]interface{})[0].(string), nil
}
