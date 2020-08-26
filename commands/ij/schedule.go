//+build ij

package ij

import "github.com/ninjawarrior1337/hanamaru-go/framework"

var Schedule = &framework.Command{
	Name:               "schedule",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Display's Mr.Z's version of the school schedule",
	Exec: func(ctx *framework.Context) error {
		ctx.Reply("https://pbs.twimg.com/media/EedUxuhU0AciSoo.png")
		return nil
	},
	Setup: nil,
}
