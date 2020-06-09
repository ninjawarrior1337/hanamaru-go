package info

import (
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var Avatar = &framework.Command{
	Name: "avatar",
	Exec: func(ctx *framework.Context) error {
		user, err := ctx.GetUser(0)
		if err != nil {
			return fmt.Errorf("please enter a valid user")
		}
		ctx.Reply(user.AvatarURL(""))
		return nil
	},
}
