package info

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/version"
)

var About = &framework.Command{
	Name: "about",
	Exec: func(ctx *framework.Context) error {
		avatarUrl := ctx.Hanamaru.State.User.AvatarURL("2048")
		_, _ = ctx.ReplyEmbed(&discordgo.MessageEmbed{
			URL:         "https://github.com/ninjawarrior1337/framework-go",
			Title:       "初めまして、",
			Description: `I am a bot created by Treelar#1974 built as a replacement for crocs-and-socks which has been put into the archives. I am built using Go and smaller and faster than crocs-and-socks and also comes with some helpful Japanese learning features if the JP build tag has been added.`,
			Color:       0x3399ff,
			Image:       nil,
			Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: avatarUrl},
			Video:       nil,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Version",
					Value: version.Version(),
				},
			},
			Author: &discordgo.MessageEmbedAuthor{
				URL:          "https://github.com/ninjawarrior1337/hanamaru-go",
				Name:         "Hanamaru",
				IconURL:      avatarUrl,
				ProxyIconURL: "",
			},
		})
		return nil
	},
}
