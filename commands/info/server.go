package info

import (
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var ServerInfo = &framework.Command{
	Name:               "server",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Exec: func(ctx *framework.Context) error {
		serverid, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		g, err := ctx.Hanamaru.Guild(serverid)
		if err != nil {
			return err
		}
		ctx.Reply(fmt.Sprintf("%v: %v", g.Name, g.Channels))
		return nil
	},
}
