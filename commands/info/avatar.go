package info

import (
	"fmt"
	"hanamaru/hanamaru"
)

var Avatar = &hanamaru.Command{
	Name: "avatar",
	Exec: func(ctx *hanamaru.Context) error {
		user, err := ctx.GetUser(0)
		if err != nil {
			return fmt.Errorf("please enter a valid user")
		}
		ctx.Reply(user.AvatarURL(""))
		return nil
	},
}