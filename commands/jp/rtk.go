//go:build jp

package jp

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util/jp"
)

var rtkDB *jp.RtkDatabase

var Rtk = &framework.Command{
	Name:               "rtk",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		search := ctx.TakeRest()
		search = strings.TrimSpace(search)
		entry, err := rtkDB.Search(search)
		if err != nil {
			return err
		}
		ctx.ReplyEmbed(constructEmbedFromEntry(entry))
		return nil
	},
	Setup: func() error {
		rtkDB = jp.NewRtkDatabase()
		return nil
	},
}

func constructEmbedFromEntry(entry jp.RtkEntry) *discordgo.MessageEmbed {
	e := &discordgo.MessageEmbed{
		URL:         fmt.Sprintf("https://hochanh.github.io/rtk/%v/index.html", string(entry.Kanji)),
		Title:       string(entry.Kanji),
		Description: entry.Story,
		Color:       0x007ade,
		Image:       nil,
		Thumbnail:   nil,
		Author:      nil,
		Fields:      []*discordgo.MessageEmbedField{},
	}
	for i, story := range entry.Koohi {
		e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("Koohi Story %v", i+1),
			Value:  story,
			Inline: false,
		})
	}
	return e
}
