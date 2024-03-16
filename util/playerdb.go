package util

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    Data   `json:"data"`
}

type Data struct {
	Player Player `json:"player"`
}

type Player struct {
	Username string                 `json:"username"`
	Id       string                 `json:"id"`
	Avatar   string                 `json:"avatar"`
	Meta     map[string]interface{} `json:"meta"`
}

type MinecraftRenderType string

const (
	Avatars MinecraftRenderType = "avatars"
	Head                        = "head"
	Body                        = "body"
	Skins                       = "skins"
)

func LookupMinecraft(playerName string) (*Player, error) {
	resp, err := http.Get(fmt.Sprintf("https://playerdb.co/api/player/minecraft/%v", playerName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r = new(Response)
	json.NewDecoder(resp.Body).Decode(&r)
	return &r.Data.Player, nil
}

func GetMinecraftSkin(player *Player, renderType MinecraftRenderType) (image.Image, error) {
	var urlBase string
	switch renderType {
	case Avatars:
		urlBase = "https://crafatar.com/avatars/%v"
	case Body:
		urlBase = "https://crafatar.com/renders/body/%v?overlay"
	case Head:
		urlBase = "https://crafatar.com/renders/head/%v"
	case Skins:
		urlBase = "https://crafatar.com/skins/%v"
	}
	resp, err := http.Get(fmt.Sprintf(urlBase, player.Id))
	if err != nil || resp.StatusCode != 200 {
		if resp != nil && resp.StatusCode != 200 {
			return nil, fmt.Errorf("failed to get skin with UUID: %v", player.Id)
		}
		return nil, err
	}
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	return img, err
}
