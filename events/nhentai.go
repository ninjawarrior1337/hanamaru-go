package events

import (
	"github.com/bwmarrin/discordgo"
	"hanamaru/util"
	"regexp"
	"strconv"
	"strings"
)

var nhr = regexp.MustCompile(`(\d{6})`)

var Nhentai = func(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	matches := nhr.FindStringSubmatch(m.Content)
	if len(matches) == 0 {
		return
	}
	matchInt, _ := strconv.Atoi(matches[0])

	n, err := util.ParseNhentai(matchInt)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, strconv.Itoa(matchInt)+": Not Found")
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, ConstructEmbed(n))
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
