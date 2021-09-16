//go:build jp

package jp

import (
	"fmt"
	"strings"

	"github.com/ninjawarrior1337/hanamaru-go/framework"

	"github.com/ninjawarrior1337/hanamaru-go/util/jp"
)

var Romaji = &framework.Command{
	Name:               "roma",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		input := strings.TrimLeft(ctx.TakeRest(), " ")
		if input == "" {
			return fmt.Errorf("please input a string to turn into romaji")
		}
		output := jp.ParseJapanese(input)
		ctx.Reply(output)
		return nil
	},
}
