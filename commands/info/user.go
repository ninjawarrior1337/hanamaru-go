package info

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var UserInfo = &framework.Command{
	Name:               "user",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Exec: func(ctx *framework.Context) error {
		//ctx.Session.user
		return nil
	},
}

func constructEmbed(member *discordgo.Member) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "",
		Type:        "",
		Title:       "",
		Description: "",
		Timestamp:   "",
		Color:       0,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
		Fields: []*discordgo.MessageEmbedField{
			{"Full Name", fmt.Sprintf("%v#%v", member.User.Username, member.User.Discriminator), true},
			{"Nickname", member.Nick, true},
			//{"Account Created", member.},
		},
	}
}
