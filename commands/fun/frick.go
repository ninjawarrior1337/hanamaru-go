package fun

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var Frick = &framework.Command{
	Name:               "frick",
	PermissionRequired: discordgo.PermissionManageChannels | discordgo.PermissionManageServer,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		tmpCh, err := ctx.Hanamaru.GuildChannelCreate(ctx.GuildID, "this is death", discordgo.ChannelTypeGuildVoice)
		if err != nil {
			return errors.New("failed to create temporary channel, check the permissions given to the bot")
		}
		fromVC, err := ctx.GetSenderVoiceChannel()
		if err != nil {
			return nil
		}
		guild, err := ctx.Hanamaru.Guild(ctx.GuildID)
		if err != nil {
			return nil
		}
		for _, s := range guild.VoiceStates {
			if s.ChannelID == fromVC.ID {
				ctx.Hanamaru.GuildMemberMove(ctx.GuildID, s.UserID, &tmpCh.ID)
			}
		}
		ctx.Hanamaru.ChannelDelete(tmpCh.ID)
		ctx.Reply("Done")
		return nil
	},
}
