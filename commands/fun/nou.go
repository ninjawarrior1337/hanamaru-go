package fun

import (
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var NoU = &framework.Command{
	Name:               "nou",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		ctx.Reply("https://i.imgur.com/3WDcYbV.png")
		return nil
	},
}
