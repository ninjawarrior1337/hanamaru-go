package info

import (
	"fmt"
	"math"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
)

var WhatAnimeIsThat = &framework.Command{
	Name:               "wait",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Uses trace.moe to look up what anime the picture just sent is from",
	Exec: func(ctx *framework.Context) error {
		img, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		ta, err := util.TraceAnime(img)
		if err != nil {
			return err
		}

		if len(ta.Result) > 0 {
			_, err = ctx.ReplyEmbed(waitEmbed(ta.Result[0]))
			return err
		}
		return nil
	},
	Setup: nil,
}

func waitEmbed(doc util.TraceMoeResult) *discordgo.MessageEmbed {
	al, _ := util.GetAnimeInfoFromID(doc.Anilist)
	return &discordgo.MessageEmbed{
		Title:     "I think that this comes from " + al.Title.English,
		URL:       al.SiteURL,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: al.CoverImage.Large},
		Color:     0x02a9ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Appears at",
				Value:  fmt.Sprintf("%02.f:%02.f", doc.From/60, math.Mod(doc.From, 60)),
				Inline: true,
			},
			{
				Name:   "Episode",
				Value:  strconv.Itoa(doc.Episode),
				Inline: true,
			},
			{
				Name:  "Similarity",
				Value: fmt.Sprintf("%02.f%%", doc.Similarity*100),
			},
		},
	}
}
