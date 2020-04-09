package hanamaru

import (
	"fmt"
	"image"
	"net/http"
	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
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
	var imgUrl string

	if len(c.Message.Attachments) <= 0 {
		if len(c.Args) == 0 {
			return nil, fmt.Errorf("this doesn't contain an image")
		} else {
			if _, err := url.Parse(c.Args[0]); err != nil {
				return nil, fmt.Errorf("this doesn't contain an image")
			}
			imgUrl = c.Args[0]
		}
	}

	imgUrl = c.Message.Attachments[idx].URL

	resp, err := http.Get(imgUrl)
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

func (c *Context) GetUser(idx int) (*discordgo.User, error) {
	if len(c.Mentions) <= 0 {
		if len(c.Args) == 0 {
			return nil, fmt.Errorf("this doesn't contain any mentions")
		} else {
			usr, err := c.Session.User(c.Args[idx])
			if err != nil {
				return nil, fmt.Errorf("invalid user: %v", err)
			}
			return usr, nil
		}
	}
	return c.Mentions[idx], nil
}

func (c *Context) GetMember(idx int) (*discordgo.Member, error) {
	if len(c.Mentions) <= 0 {
		return nil, fmt.Errorf("this doesn't contain any mentions")
	}
	return c.Session.GuildMember(c.GuildID, c.Mentions[idx].ID)
}
