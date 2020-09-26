package util

import "testing"

func TestAnilist_GetAnimeInfo(t *testing.T) {
	SkipCI(t)
	m, err := GetAnimeInfoFromTitle("Love Live Sunshine")
	if err != nil {
		t.Error(err)
		return
	}

	if m.Title.English != "Love Live! Sunshine!!" {
		t.Error()
		return
	}
}
