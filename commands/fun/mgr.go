package fun

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var Migrate = &framework.Command{
	Name:               "mgr",
	PermissionRequired: discordgo.PermissionManageChannels | discordgo.PermissionManageServer,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		toVC, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		_, err = ctx.Guild(toVC)
		if err != nil {
			return errors.New("please use a guild id that exists")
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
				ctx.GuildMemberMove(ctx.GuildID, s.UserID, toVC)
			}
		}
		return nil
	},
}
