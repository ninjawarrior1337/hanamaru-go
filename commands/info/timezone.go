package info

import (
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
	"time"
)

var Timezone = &framework.Command{
	Name:               "tz",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               `Converts the current time to another timezone. <time|now fmt. hh:mmAM/PM> <dest zone> <src zone>`,
	Exec: func(ctx *framework.Context) error {
		timeStr, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		destTzStr, err := ctx.GetArgIndex(1)
		if err != nil {
			return err
		}
		srcTzStr := ctx.GetArgIndexDefault(2, "America/Los_Angeles")
		t, err := util.ConvertTimezone(timeStr, srcTzStr, destTzStr)
		if err != nil {
			return err
		}
		ctx.Reply(t.Format(time.Stamp))
		return nil
	},
}
