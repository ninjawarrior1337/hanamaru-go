package voice

import "testing"

var y *Ytdl

func init() {
	y, _ = NewYTDL("")
}

func TestYtdl_GetVideoInfo(t *testing.T) {
	v, err := y.GetVideoInfo("https://www.youtube.com/watch?v=XQsMmtC91b4")
	if err != nil {
		t.Error(err)
	}
	if v.Title != "Bossfight - Commando Steve" {
		t.Error()
	}

	_, err = y.GetVideoInfo("https://www.youtube.com/watch?v=nothiscantexist")
	if err == nil {
		t.Error(err)
	}
}

func TestYtdl_GetVideoInfoSoundcloud(t *testing.T) {
	v, err := y.GetVideoInfo("https://soundcloud.com/uiceheidd/desire")
	if err != nil {
		t.Error(nil)
	}
	if v.Title != "Desire" {
		t.Error()
	}
}
