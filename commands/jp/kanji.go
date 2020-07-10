package jp

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util/jisho"
	"strconv"
	"strings"
	"unicode/utf8"
)

var KanjiCmd = &framework.Command{
	Name:               "kanji",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		kanjiStr, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		if utf8.RuneCountInString(kanjiStr) > 1 {
			return errors.New("kanji are always one character")
		}
		k, err := j.SearchForKanji(kanjiStr)
		if err != nil {
			return err
		}
		_, err = ctx.ReplyEmbed(TransformKanjiToEmbed(k))
		return err
	},
}

func TransformKanjiToEmbed(k *jisho.Kanji) *discordgo.MessageEmbed {
	var fl = []*discordgo.MessageEmbedField{
		{
			Name:  "Taught In",
			Value: strings.Title(k.TaughtIn),
		},
		{
			Name:  "JLPT",
			Value: string(k.JLPT),
		},
		{
			Name:  "Newspaper Frequency",
			Value: strconv.Itoa(k.NewspaperFreqRank),
		},
		{
			Name:  "Stroke Count",
			Value: strconv.Itoa(k.StrokeCount),
		},
		{
			Name:  "Kunyomi",
			Value: strings.Join(k.Kunyomi, ", "),
		},
		{
			Name:  "Onyomi",
			Value: strings.Join(k.Onyomi, ", "),
		},
		{
			Name:  "Radical",
			Value: k.Radical.Symbol + ": " + k.Radical.Meaning,
		},
		{
			Name:  "Parts",
			Value: strings.Join(k.Parts, ", "),
		},
	}
	for _, f := range fl {
		f.Inline = true
	}
	return &discordgo.MessageEmbed{
		URL:         k.JishoUri,
		Title:       string(k.Rune),
		Description: fmt.Sprintf("**%s**", k.Meaning),
		Color:       0x56D926,
		Footer:      nil,
		Image: &discordgo.MessageEmbedImage{
			URL: k.StrokeOrderDiagram,
		},
		Thumbnail: nil,
		Author:    nil,
		Fields:    fl,
	}
}
