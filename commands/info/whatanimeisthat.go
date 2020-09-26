package info

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
	"math"
	"strconv"
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

		if len(ta.Docs) > 0 {
			_, err = ctx.ReplyEmbed(waitEmbed(ta.Docs[0]))
			return err
		}
		return nil
	},
	Setup: nil,
}

func waitEmbed(doc util.TraceMoeDocs) *discordgo.MessageEmbed {
	al, _ := util.GetAnimeInfoFromID(doc.AnilistID)
	return &discordgo.MessageEmbed{
		Title:     "I think that this comes from " + doc.TitleEnglish,
		URL:       al.SiteURL,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: al.CoverImage.Large},
		Color:     0x02a9ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Appears at",
				Value:  fmt.Sprintf("%02.f:%02.f", doc.At/60, math.Mod(doc.At, 60)),
				Inline: true,
			},
			{
				Name:   "Episode",
				Value:  strconv.Itoa(doc.Episode),
				Inline: true,
			},
			{
				Name:  "Certainty",
				Value: fmt.Sprintf("%02.f%%", doc.Similarity*100),
			},
		},
	}
}
