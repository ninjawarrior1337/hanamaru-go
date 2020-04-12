// +build jp

package jp

import (
	"fmt"
	"hanamaru/hanamaru"
	"hanamaru/util"
)

var Romaji = &hanamaru.Command{
	Name:               "roma",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		input, err := ctx.GetArgIndex(0)
		if err != nil {
			return fmt.Errorf("please input a string to turn into romaji")
		}
		output := util.ParseJapanese(input)
		ctx.Reply(output)
		return nil
	},
}
