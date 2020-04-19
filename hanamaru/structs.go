package hanamaru

import (
	"fmt"
	"hanamaru/hanamaru/voice"
	"image"
	"io"
	"net/http"
	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
)

type Command struct {
	Name               string
	PermissionRequired int
	OwnerOnly          bool
	Exec               func(ctx *Context) error
}

type Context struct {
	*discordgo.Session
	*discordgo.MessageCreate
	Args         []string
	VoiceContext *voice.Context
}

func (c *Context) Reply(m string) (*discordgo.Message, error) {
	return c.ChannelMessageSend(c.ChannelID, m)
}

func (c *Context) ReplyFile(name string, r io.Reader) (*discordgo.Message, error) {
	return c.ChannelFileSend(c.ChannelID, name, r)
}

func (c *Context) GetImage(idx uint) (*gg.Context, error) {
	var imgUrl string

	if len(c.Message.Attachments) <= 0 {
		if len(c.Args) == 0 {
			return nil, fmt.Errorf("this doesn't contain an image")
		} else {
			if _, err := url.Parse(c.Args[0]); err != nil {
				return nil, fmt.Errorf("this invalid image: %v", c.Args[0])
			}
			imgUrl, c.Args = c.Args[0], c.Args[1:]
		}
	} else {
		imgUrl = c.Message.Attachments[idx].URL
	}

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

func (c *Context) GetVoiceChannnel() (*discordgo.Channel, error) {
	guild, err := c.Guild(c.GuildID)
	if err != nil {
		return nil, fmt.Errorf("you can only use this command in a guild")
	}
	for _, state := range guild.VoiceStates {
		if state.UserID == c.Author.ID {
			channel, _ := c.State.Channel(state.ChannelID)

			return channel, nil
		}
	}

	return nil, fmt.Errorf("please join a vc before using this command")
}

func (c *Context) GetArgIndex(idx int) (string, error) {
	if idx > len(c.Args)-1 {
		return "", fmt.Errorf("failed to get arg with index %v", idx)
	}
	return c.Args[idx], nil
}
