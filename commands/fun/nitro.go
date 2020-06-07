package fun

import (
	"github.com/bwmarrin/discordgo"
	"hanamaru/hanamaru"
)

var Nitro = &hanamaru.Command{
	Name:               "nitro",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *hanamaru.Context) error {
		channel, err := ctx.Channel(ctx.ChannelID)
		if err != nil {
			return err
		}
		if channel.Type != discordgo.ChannelTypeDM {
			if channel.Type != discordgo.ChannelTypeGroupDM {
				ctx.ChannelMessageDelete(ctx.ChannelID, ctx.Message.ID)
				ctx.ReplyEmbed(constructNitroEmbed())
			}
		} else {
			ctx.ReplyEmbed(constructNitroEmbed())
		}
		return nil
	},
}

func constructNitroEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Description: "[Discord Nitro](https://discordapp.com/nitro) is **required** to view this message.",
		Color:       5267072,
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: "https://cdn.discordapp.com/attachments/194167041685454848/272617748876492800/be14b7a8e0090fbb48135450ff17a62f.png"},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Discord Nitro Message",
			IconURL: "https://cdn.discordapp.com/emojis/264287569687216129.png",
		},
	}
}
