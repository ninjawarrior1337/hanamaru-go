package fun

import "hanamaru/hanamaru"

var NoU = &hanamaru.Command{
	Name:               "nou",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *hanamaru.Context) error {
		ctx.Reply("https://i.imgur.com/3WDcYbV.png")
		return nil
	},
}
