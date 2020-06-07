package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"hanamaru/util"
	"regexp"
	"strconv"
	"strings"
)

var nhr = regexp.MustCompile(`^(\d{6})$`)

var Nhentai = func(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || len(m.Mentions) > 0 {
		return
	}
	//TODO: Have unified way to get the active prefix of the bot
	if strings.HasPrefix(m.Content, "!") {
		return
	}
	if channel, _ := s.Channel(m.ChannelID); channel != nil && !channel.NSFW {
		return
	}
	matches, err := ParseStringWithSixDigits(m.Content)
	if err != nil {
		return
	}

	for _, match := range matches {
		n, err := util.ParseNhentai(match)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, strconv.Itoa(match)+": Not Found")
		}
		s.ChannelMessageSendEmbed(m.ChannelID, ConstructEmbed(n))
	}
}

func ParseStringWithSixDigits(msg string) ([]int, error) {
	msgs := strings.Split(msg, " ")
	var intMatches []int
	for _, msg := range msgs {
		matches := nhr.FindAllStringSubmatch(msg, -1)
		if len(matches) == 0 {
			continue
		}
		for _, match := range matches {
			if len(intMatches) >= 3 {
				return intMatches, nil
			}
			matchInt, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}
			intMatches = append(intMatches, matchInt)
		}
	}
	if len(intMatches) == 0 {
		return nil, fmt.Errorf("no matches")
	}

	return intMatches, nil
}

func ConstructEmbed(n util.NHentai) *discordgo.MessageEmbed {
	var fields []*discordgo.MessageEmbedField
	for name, tags := range n.Tags {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   name,
			Value:  strings.Join(tags, ", "),
			Inline: false,
		})
	}

	return &discordgo.MessageEmbed{
		URL:       n.URL,
		Type:      "rich",
		Title:     n.Title,
		Color:     0xed2553,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: n.CoverArt},
		Fields:    fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: strconv.Itoa(n.PageCount) + " pages",
		},
	}

}
