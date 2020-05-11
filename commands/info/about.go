package info

import (
	"github.com/bwmarrin/discordgo"
	"hanamaru/hanamaru"
	"os"
)

var About = &hanamaru.Command{
	Name: "about",
	Exec: func(ctx *hanamaru.Context) error {
		_, _ = ctx.ChannelMessageSendEmbed(ctx.ChannelID, &discordgo.MessageEmbed{
			URL:         "https://github.com/ninjawarrior1337/hanamaru-go",
			Type:        "",
			Title:       "初めまして、",
			Description: `I am a bot created by Treelar#1974 built as a replacement for crocs-and-socks which has been put into the archives. I am built using Go and smaller and faster than crocs-and-socks and also comes with some cool *Dogen-like* features if the JP build tag has been added.`,
			Color:       0x3399ff,
			Image:       nil,
			Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: "https://cdn.discordapp.com/avatars/405165920563232778/fd1727a732d6605c49f0b729c4bb6e89.png"},
			Video:       nil,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: "Build Date",
					Value: func() string {
						if bd := os.Getenv("BUILD_DATE"); bd == "" {
							return "develop"
						} else {
							return bd
						}
					}(),
				},
			},
			Author: &discordgo.MessageEmbedAuthor{
				URL:          "https://github.com/ninjawarrior1337/hanamaru-go",
				Name:         "Hanamaru",
				IconURL:      "https://cdn.discordapp.com/avatars/405165920563232778/fd1727a732d6605c49f0b729c4bb6e89.png",
				ProxyIconURL: "",
			},
		})
		return nil
	},
}
