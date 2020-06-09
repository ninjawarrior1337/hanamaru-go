package debug

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var MsgInfo = &framework.Command{
	Name:               "msginfo",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		msg, err := ctx.GetPreviousMessage()
		if err != nil {
			return err
		}
		ctx.Reply("Message info has been sent to the console")
		spew.Dump(msg)
		return nil
	},
}
