package jp

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util/jisho"
	"strings"
)

var JishoCmd = &framework.Command{
	Name:               "jisho",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Searches jisho and displays the first result",
	Exec: func(ctx *framework.Context) error {
		search := ctx.TakeRest()[1:]
		sr, err := j.SearchKeyword(search)
		if err != nil {
			return err
		}
		ctx.ReplyEmbed(JishoResponseAsEmbed(sr))
		return nil
	},
	Setup: nil,
}

func JishoResponseAsEmbed(sr *jisho.SearchResp) *discordgo.MessageEmbed {
	d := sr.Data[0]
	return &discordgo.MessageEmbed{
		URL:         fmt.Sprintf("https://jisho.org/search/%s", sr.Data[0].Slug),
		Title:       d.Slug + fmt.Sprintf(" (%s)", d.Japanese[0].Reading),
		Description: strings.Join(d.Senses[0].EnglishDefinitions, ", "),
		Timestamp:   "",
		Color:       0x56D926,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Tags",
				Value:  strings.Join(d.Tags, ", "),
				Inline: true,
			},
			{
				Name: "Common?",
				Value: func() string {
					if d.IsCommon {
						return "Yes"
					} else {
						return "No"
					}
				}(),
				Inline: true,
			},
			{
				Name:   "Part of Speech",
				Value:  strings.Join(d.Senses[0].PartsOfSpeech, ", "),
				Inline: true,
			},
			{
				Name:  "JLPT",
				Value: strings.ToUpper(strings.Join(d.Jlpt, ", ")),
			},
		},
	}
}
