package fun

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"hanamaru/hanamaru"
)

var Frick = &hanamaru.Command{
	Name:               "frick",
	PermissionRequired: discordgo.PermissionManageChannels | discordgo.PermissionManageServer,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *hanamaru.Context) error {
		tmpCh, err := ctx.GuildChannelCreate(ctx.GuildID, "this is death", discordgo.ChannelTypeGuildVoice)
		if err != nil {
			return errors.New("failed to create temporary channel, check the permissions given to the bot")
		}
		fromVC, err := ctx.GetSenderVoiceChannel()
		if err != nil {
			return nil
		}
		guild, err := ctx.Guild(ctx.GuildID)
		if err != nil {
			return nil
		}
		for _, s := range guild.VoiceStates {
			if s.ChannelID == fromVC.ID {
				ctx.GuildMemberMove(ctx.GuildID, s.UserID, tmpCh.ID)
			}
		}
		ctx.ChannelDelete(tmpCh.ID)
		ctx.Reply("Done")
		return nil
	},
}
