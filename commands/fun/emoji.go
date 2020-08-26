package fun

import (
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var Emoji = &framework.Command{
	Name: "emoji",
	Exec: func(ctx *framework.Context) error {
		name, err := ctx.GetArgIndex(0)
		id, err := ctx.GetArgIndex(1)
		if err != nil {
			return err
		}
		ctx.Reply(fmt.Sprintf("<:%s:%s>", name, id))
		return nil
	},
}
