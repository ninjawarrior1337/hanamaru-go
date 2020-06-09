package image

import (
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

var Braile = &framework.Command{
	Name:               "braile",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {

		return nil
	},
}
