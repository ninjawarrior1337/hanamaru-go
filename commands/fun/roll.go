package fun

import (
	"fmt"
	"hanamaru/hanamaru"
	"math/rand"
	"strconv"
)

var Roll = &hanamaru.Command{
	Name:               "roll",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *hanamaru.Context) error {
		rollStr := ctx.GetArgIndexDefault(0, "100")
		rollInt, err := strconv.Atoi(rollStr)
		if err != nil {
			return fmt.Errorf("%v is not a number", rollStr)
		}
		ctx.Reply(fmt.Sprintf("%v rolls %v point(s)!", ctx.Author.Username, rand.Intn(rollInt)+1))
		return nil
	},
}
