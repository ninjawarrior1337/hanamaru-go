package info

import "hanamaru/hanamaru"

var ServerInfo = &hanamaru.Command{
	Name:               "server",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Exec: func(ctx *hanamaru.Context) error {
		return nil
	},
}
