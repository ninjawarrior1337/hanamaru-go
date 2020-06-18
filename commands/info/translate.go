package info

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
)

var Translate = &framework.Command{
	Name:               "translate",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Translates text: <text> <src lang (def. auto)> <target lang (def. en)>",
	Exec: func(ctx *framework.Context) error {
		text, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		srcLang := ctx.GetArgIndexDefault(1, "")
		destLang := ctx.GetArgIndexDefault(2, "en")
		tr, err := util.Translate(text, srcLang, destLang)
		ctx.ReplyEmbed(GenTranslationEmbed(text, tr))
		return nil
	},
}

func GenTranslationEmbed(orig, trans string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0x3399ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Original",
				Value: orig,
			},
			{
				Name:  "Translated",
				Value: trans,
			},
		},
	}
}
