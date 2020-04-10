package music

import (
	"fmt"
	"hanamaru/hanamaru"
)

var Leave = &hanamaru.Command{
	Name:               "leave",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		if val, ok := ctx.VoiceContext.VCs[ctx.GuildID]; !ok {
			return fmt.Errorf("cannot disconnect when im not connected")
		} else {
			return val.Disconnect()
		}
	},
}
