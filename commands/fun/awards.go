package fun

import (
	"fmt"
	"strings"

	"github.com/ninjawarrior1337/hanamaru-go/events"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
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
				akv := events.GetUserAwards(ctx.Hanamaru.Db, ctx.Author.ID)
				ctx.Reply(formatAwardKV(akv))
			}
		case "leaderboard":
			{
				ctx.Reply(generateLeaderboard(ctx.Hanamaru))
			}
		}

		return nil

	},
	Setup: nil,
}

func generateLeaderboard(h *framework.Hanamaru) string {
	type LeaderData struct {
		Count int
		Uid   string
	}
	leaderBoard := make(map[string]LeaderData)
	ai := events.GetAllAwardsInfo(h.Db)
	for id, akv := range ai {
		for award, num := range akv {
			awardStr := string(award)
			// Make sure values about to be compared actually exist
			if _, ok := leaderBoard[awardStr]; !ok {
				leaderBoard[awardStr] = LeaderData{}
			}
			if num > leaderBoard[awardStr].Count {
				leaderBoard[awardStr] = LeaderData{
					Count: num,
					Uid:   id,
				}
			}
		}
	}
	finalString := "Leaderboard: \n"

	for k, v := range leaderBoard {
		user, err := h.User(v.Uid)
		if err != nil {
			return "Failed to generate leaderboard: " + err.Error()
		}
		finalString += fmt.Sprintf("Most %v: %v with %d award(s)\n", strings.Split(k, "_")[0], user.String(), v.Count)
	}
	return finalString
}

func formatAwardKV(akv events.AwardKV) string {
	var output = "Here are your awards stats: \n"
	for k, v := range akv {
		name, _ := k.Name()
		output += fmt.Sprintf("%v award: %d \n", name, v)
	}
	return output
}
