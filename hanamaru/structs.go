package hanamaru

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
	"image"
	"net/http"
)

type Command struct {
	Name               string
	PermissionRequired int
	Exec               func(ctx *Context) error
}

type Context struct {
	*discordgo.Session
	*discordgo.MessageCreate
	Args []string
}

func (c *Context) Reply(m string) (*discordgo.Message, error) {
	return c.ChannelMessageSend(c.ChannelID, m)
}

func (c *Context) GetImage(idx uint) (*gg.Context, error) {
	if len(c.Message.Attachments) <= 0 {
		return nil, fmt.Errorf("this doesn't contain an image")
	}

	resp, err := http.Get(c.Message.Attachments[idx].URL)
	if err != nil {
		return nil, fmt.Errorf("failed to download image from discord: %v", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image sent: %v", img)
	}

	return gg.NewContextForImage(img), nil
}
