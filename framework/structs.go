package framework

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type EventListener struct {
	// Name of the event listener.
	Name string
	// HandlerConstructor is a function that depends on a Hanamaru pointer and must return an event listener
	HandlerConstructor func(h *Hanamaru) interface{}
}

type Command struct {
	Name               string
	PermissionRequired int64
	OwnerOnly          bool
	Help               string
	Exec               func(ctx *Context) error
	Setup              func() error
}

type Context struct {
	Hanamaru *Hanamaru
	Command  *Command
	*discordgo.MessageCreate
	Args []string
}

func NewContext(h *Hanamaru, cmd *Command, m *discordgo.MessageCreate) *Context {
	ctx := &Context{
		Hanamaru:      h,
		Command:       cmd,
		MessageCreate: m,
		Args:          []string{},
	}
	argsString := strings.TrimPrefix(m.Content, h.Prefix+cmd.Name)
	ctx.Args = ParseArgs(argsString)
	return ctx
}

func (c *Context) Reply(m string) (*discordgo.Message, error) {
	if m == "" {
		return c.Hanamaru.ChannelMessageSendComplex(c.ChannelID, &discordgo.MessageSend{
			Content:   "ERROR...idk",
			Reference: c.Reference(),
		})
	}
	return c.Hanamaru.ChannelMessageSendComplex(c.ChannelID, &discordgo.MessageSend{
		Content:   m,
		Reference: c.Reference(),
	})
}

func (c *Context) ReplyEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	var validFields []*discordgo.MessageEmbedField
	for _, f := range embed.Fields {
		if f.Name != "" && f.Value != "" {
			validFields = append(validFields, f)
		}
	}

	embed.Fields = validFields

	return c.Hanamaru.ChannelMessageSendComplex(c.ChannelID, &discordgo.MessageSend{
		Embed:     embed,
		Reference: c.Reference(),
	})
}

func (c *Context) ReplyFile(name string, r io.Reader) (*discordgo.Message, error) {
	return c.Hanamaru.ChannelMessageSendComplex(c.ChannelID, &discordgo.MessageSend{
		File: &discordgo.File{
			Name:   name,
			Reader: r,
		},
		Reference: c.Reference(),
	})
}

// ReplyPNGImg Sends img into the chat where the command is executed. Name excludes the extension
func (c *Context) ReplyPNGImg(img image.Image, name string) (*discordgo.Message, error) {
	var pngBuf = new(bytes.Buffer)
	err := png.Encode(pngBuf, img)
	if err != nil {
		return nil, fmt.Errorf("failed to encode png image, please report to Treelar#1974: %v", err)
	}
	return c.ReplyFile(name+".png", pngBuf)
}

// This is name without extension btw, the following function will add it by itself
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

	if len(c.Message.Attachments) > 0 {
		imgUrl = c.Message.Attachments[idx].URL
		// if ends here
	} else {
		if len(c.Args) != 0 {
			if _, err := url.Parse(c.Args[idx]); err != nil {
				imgUrl = ""
			}
			imgUrl, c.Args = c.Args[0], c.Args[1:]
			// if ends here
		} else {
			msg, err := c.GetPreviousMessage()
			if err != nil {
				imgUrl = ""
				// if ends here
			} else if len(msg.Attachments) > 0 {
				imgUrl = msg.Attachments[idx].URL
				// if ends here
			} else if len(msg.Embeds) > 0 {
				imgUrl = msg.Embeds[0].URL
			}
		}
	}

	if imgUrl == "" {
		return nil, fmt.Errorf("unable to determine which image to use")
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
			usr, err := c.Hanamaru.User(c.Args[idx])
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
	return c.Hanamaru.GuildMember(c.GuildID, c.Mentions[idx].ID)
}

func (c *Context) GetSenderVoiceChannel() (*discordgo.Channel, error) {
	guild, err := c.Hanamaru.Guild(c.GuildID)
	if err != nil {
		return nil, fmt.Errorf("you can only use this command in a guild")
	}
	for _, state := range guild.VoiceStates {
		if state.UserID == c.Author.ID {
			channel, _ := c.Hanamaru.State.Channel(state.ChannelID)

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

func (c *Context) TakeRest() string {
	msgArgs := strings.TrimPrefix(c.Message.Content, c.Hanamaru.Prefix+c.Command.Name)
	return strings.ReplaceAll(msgArgs, "```", "")
}

func (c *Context) GetPreviousMessage() (*discordgo.Message, error) {
	msgs, err := c.Hanamaru.ChannelMessages(c.ChannelID, 1, c.Message.ID, "", "")
	if err != nil {
		return nil, fmt.Errorf("no messages found before previously executed command")
	}
	return msgs[0], nil
}
