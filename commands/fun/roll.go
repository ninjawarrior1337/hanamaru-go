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
		rollInt, err := strconv.Atoi(rollStr)
		if err != nil {
			return fmt.Errorf("%v is not a number", rollStr)
		}
		ctx.Reply(fmt.Sprintf("%v rolls %v point(s)!", ctx.Author.Username, rand.Intn(rollInt)+1))
		return nil
	},
}
