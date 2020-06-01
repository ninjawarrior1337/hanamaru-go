package events

import (
	"github.com/bwmarrin/discordgo"
	"hanamaru/util"
	"regexp"
	"strconv"
	"strings"
)

var animeRegex = regexp.MustCompile(`{(.*)}`)

var Roboragi = func(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if matches := animeRegex.FindAllStringSubmatch(m.Content, -1); len(matches) > 0 {
		media, err := util.GetAnimeInfo(matches[0][1])
		if err != nil {
			return
		}
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, RoboragiEmbed(media))
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		}
	}
}

func RoboragiEmbed(media util.ALMedia) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         media.SiteURL,
		Type:        "rich",
		Title:       media.Title.English,
		Description: strings.TrimSpace(strings.ReplaceAll(media.Description, "<br>", "")),
		Timestamp:   "",
		Color:       0,
		Footer:      &discordgo.MessageEmbedFooter{Text: "Status: " + strings.Title(strings.ToLower(media.Status))},
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: media.CoverImage.Large},
		Provider:    nil,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Episode Count",
				Value:  strconv.Itoa(media.Episodes),
				Inline: true,
			},
			{
				Name:   "Episode Duration",
				Value:  strconv.Itoa(media.Duration),
				Inline: true,
			},
			{
				Name:   "Score",
				Value:  strconv.Itoa(media.AverageScore),
				Inline: true,
			},
			{
				Name:   "Season",
				Value:  strconv.Itoa(media.SeasonYear),
				Inline: true,
			},
			{
				Name:   "Genres",
				Value:  strings.Join(media.Genres, ", "),
				Inline: true,
			},
		},
	}
}
