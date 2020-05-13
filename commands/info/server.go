package info

import (
	"fmt"
	"hanamaru/hanamaru"
)

var ServerInfo = &hanamaru.Command{
	Name:               "server",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Exec: func(ctx *hanamaru.Context) error {
		serverid, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		g, err := ctx.Guild(serverid)
		if err != nil {
			return err
		}
		ctx.Reply(fmt.Sprintf("%v: %v", g.Name, g.Channels))
		return nil
	},
}
