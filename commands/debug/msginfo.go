package debug

import (
	"github.com/davecgh/go-spew/spew"
	"hanamaru/hanamaru"
)

var MsgInfo = &hanamaru.Command{
	Name:               "msginfo",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		msg, err := ctx.GetPreviousMessage()
		if err != nil {
			return err
		}
		ctx.Reply("Message info has been sent to the console")
		spew.Dump(msg)
		return nil
	},
}
