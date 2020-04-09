package image

import "hanamaru/hanamaru"

var Braile = &hanamaru.Command{
	Name:               "braile",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {

		return nil
	},
}
