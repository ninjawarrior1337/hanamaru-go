package fun

import (
	"errors"
	"fmt"

	"github.com/ninjawarrior1337/hanamaru-go/db"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var AwardsCommand = &framework.Command{
	Name:               "awards",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Displays info for awards given out",
	Exec: func(ctx *framework.Context) error {
		arg := ctx.GetArgIndexDefault(0, "get")

		switch arg {
		case "get":
			{
				akv, err := db.QueryAwardCounts(ctx.Hanamaru.Db, ctx.GuildID, ctx.Author.ID)
				if err != nil {
					return err
				}
				ctx.Reply(formatAwardKV(akv))
			}
		case "leaderboard":
			{
				lbString, err := generateLeaderboard(ctx.Hanamaru, ctx.GuildID)
				if err != nil {
					return err
				}
				ctx.Reply(lbString)
			}
		}

		return nil

	},
	Setup: nil,
}

var ErrGenerateLeaderboard = errors.New("failed to generate leaderboard")

func generateLeaderboard(h *framework.Hanamaru, guild_id string) (string, error) {
	lb, err := db.QueryLeaderboard(h.Db, guild_id)
	if err != nil {
		return "", errors.Join(ErrGenerateLeaderboard, err)
	}

	finalString := "Leaderboard: \n"

	for _, entry := range lb {
		user, err := h.User(entry.EarnerId)
		if err != nil {
			return "", errors.Join(ErrGenerateLeaderboard, err)
		}
		finalString += fmt.Sprintf("%v: %v with %d award(s)\n", entry.AwardName, user.Username, entry.Count)
	}
	return finalString, nil
}

func formatAwardKV(akv db.AwardKV) string {
	caser := cases.Title(language.English)
	var output = "Here are your awards stats: \n"
	for k, v := range akv {
		output += fmt.Sprintf("%v award: %d \n", caser.String(k), v)
	}
	return output
}
