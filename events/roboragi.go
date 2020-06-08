package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/util"
	"regexp"
	"strconv"
	"strings"
)

var animeRegex = regexp.MustCompile(`{(.*)}`)

var Roboragi = func(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	//TODO: Have unified way to get the active prefix of the bot
	if strings.HasPrefix(m.Content, "!") {
		return
	}
	if matches := animeRegex.FindAllStringSubmatch(m.Content, -1); len(matches) > 0 {
		media, err := util.GetAnimeInfo(matches[0][1])
		if err != nil {
			return
		}
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, roboragiEmbed(media))
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		}
	}
}

func roboragiEmbed(media util.ALMedia) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         media.SiteURL,
		Type:        "rich",
		Title:       media.Title.English,
		Description: strings.TrimSpace(strings.ReplaceAll(media.Description, "<br>", "")),
		Timestamp:   "",
		Color:       0,
		Footer:      &discordgo.MessageEmbedFooter{Text: "Status: " + strings.Title(strings.ToLower(media.Status)) + " | " + "Type: " + string(media.Type)},
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: media.CoverImage.Large},
		Provider:    nil,
		Fields:      generateFields(media),
	}
}

func generateFields(media util.ALMedia) []*discordgo.MessageEmbedField {
	var f []*discordgo.MessageEmbedField
	if media.Type == util.Anime {
		f = []*discordgo.MessageEmbedField{
			{
				Name:   "Episodes",
				Value:  strconv.Itoa(media.Episodes),
				Inline: true,
			},
			{
				Name:   "Episode Duration",
				Value:  strconv.Itoa(media.Duration) + " Minutes",
				Inline: true,
			},
			{
				Name:   "Season",
				Value:  strconv.Itoa(media.SeasonYear),
				Inline: true,
			},
		}
	} else {
		f = []*discordgo.MessageEmbedField{
			{
				Name:   "Volumes",
				Value:  strconv.Itoa(media.Volumes),
				Inline: true,
			},
			{
				Name:   "Chapters",
				Value:  strconv.Itoa(media.Chapters),
				Inline: true,
			},
		}
	}
	f = append(f,
		&discordgo.MessageEmbedField{
			Name:   "Score",
			Value:  strconv.Itoa(media.AverageScore),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Genres",
			Value:  strings.Join(media.Genres, ", "),
			Inline: true,
		})
	return f
}
