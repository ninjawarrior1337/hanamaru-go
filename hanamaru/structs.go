package hanamaru

import (
	"bytes"
	"fmt"
	"hanamaru/hanamaru/voice"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name               string
	PermissionRequired int
	OwnerOnly          bool
	Help               string
	Exec               func(ctx *Context) error
}

type Context struct {
	Hanamaru *Hanamaru
	*discordgo.Session
	*discordgo.MessageCreate
	Args         []string
	VoiceContext *voice.Context
}

func (c *Context) Reply(m string) (*discordgo.Message, error) {
	return c.ChannelMessageSend(c.ChannelID, m)
}

func (c *Context) ReplyEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return c.ChannelMessageSendEmbed(c.ChannelID, embed)
}

func (c *Context) ReplyFile(name string, r io.Reader) (*discordgo.Message, error) {
	return c.ChannelFileSend(c.ChannelID, name, r)
}

//This is name without extension btw, the following function will add it by itself
func (c *Context) ReplyPNGImg(img image.Image, name string) (*discordgo.Message, error) {
	var pngBuf = new(bytes.Buffer)
	err := png.Encode(pngBuf, img)
	if err != nil {
		return nil, fmt.Errorf("failed to encode png image, please report to Treelar#1974: %v", err)
	}
	return c.ReplyFile(name+".png", pngBuf)
}

//This is name without extension btw, the following function will add it by itself
func (c *Context) ReplyJPGImg(img image.Image, name string) (*discordgo.Message, error) {
	jpgBuf := new(bytes.Buffer)
	err := jpeg.Encode(jpgBuf, img, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to encode jpg image, please report to Treelar#1974: %v", err)
	}
	return c.ReplyFile(name+".jpg", jpgBuf)
}

func (c *Context) ReplyGIFImg(img *gif.GIF, name string) (*discordgo.Message, error) {
	gifBuf := new(bytes.Buffer)
	err := gif.EncodeAll(gifBuf, img)
	if err != nil {
		return nil, fmt.Errorf("failed to encode gif image, please report to Treelar#1974: %v", err)
	}
	return c.ReplyFile(name+".gif", gifBuf)
}

func (c *Context) GetImage(idx uint) (image.Image, error) {
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

	return img, nil
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

func (c *Context) GetSenderVoiceChannel() (*discordgo.Channel, error) {
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

func (c *Context) GetArgIndexDefault(idx int, def string) string {
	if idx > len(c.Args)-1 {
		return def
	}
	return c.Args[idx]
}

func (c *Context) GetPreviousMessage() (*discordgo.Message, error) {
	msgs, err := c.Session.ChannelMessages(c.ChannelID, 1, c.Message.ID, "", "")
	if err != nil {
		return nil, fmt.Errorf("no messages found before previously executed command")
	}
	return msgs[0], nil
}
