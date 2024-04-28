package info

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	hdb "github.com/ninjawarrior1337/hanamaru-go/db"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var Tag = &framework.Command{
	Name:               "tag",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Let's you tag messages",
	Exec: func(ctx *framework.Context) error {
		db := ctx.Hanamaru.Db
		if target := ctx.Message.ReferencedMessage; target != nil {
			name, err := ctx.GetArgIndex(0)
			if err != nil {
				return err
			}

			err = hdb.MutateAddTag(db, name, ctx.GuildID, target.ChannelID, target.ID)
			if err != nil {
				return err
			}
			ctx.Reply("Message tagged!")
		} else {
			option, err := ctx.GetArgIndex(0)
			if err != nil {
				return err
			}

			switch option {
			case "list":
				pageStr := ctx.GetArgIndexDefault(1, "1")
				page, err := strconv.Atoi(pageStr)
				if err != nil {
					return err
				}
				res, err := hdb.QueryListTags(db, ctx.GuildID, page-1)
				if err != nil {
					return err
				}

				if len(res) < 1 {
					return errors.New("no results found")
				}

				var s = "Tags: \n"
				s += strings.Join(res, "\n")
				ctx.Reply(s)

			case "delete":
				tagName, err := ctx.GetArgIndex(1)
				if err != nil {
					return err
				}
				hdb.MutateRemoveTagByName(db, ctx.GuildID, tagName)
			case "search":
				query, err := ctx.GetArgIndex(1)
				if err != nil {
					return err
				}

				res, err := hdb.QuerySearchTags(db, query, ctx.GuildID)
				if err != nil {
					return err
				}
				if len(res) < 1 {
					return errors.New("no results found")
				}

				var s = "Results: \n"
				s += strings.Join(res, "\n")
				ctx.Reply(s)
			default:
				tag, err := hdb.QueryTagByName(db, option, ctx.GuildID)
				if err != nil {
					return err
				}

				ctx.Reply(fmt.Sprintf("https://discord.com/channels/%v/%v/%v", tag.GuildID, tag.ChannelID, tag.MessageID))
			}
		}

		return nil
	},
}
