package fun

import (
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"math/rand"
	"strconv"
)

var Roll = &framework.Command{
	Name:               "roll",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		rollStr := ctx.GetArgIndexDefault(0, "100")
		rollInt64, err := strconv.ParseInt(rollStr, 10, 64)
		if err != nil {
			rollInt64 = 100
		}
		if rollInt64 <= 0 {
			ctx.Reply(fmt.Sprintf("%v is how many friends you have %s", rollInt64, ctx.Message.Author.Mention()))
			return nil
		}
		ctx.Reply(fmt.Sprintf("%v rolls %v point(s)!", ctx.Author.Username, rand.Int63n(rollInt64)+1))
		return nil
	},
}
