package image

import (
	"hanamaru/hanamaru"
	"hanamaru/util/latex"
)

var Latex = &hanamaru.Command{
	Name:               "latex",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Exec: func(ctx *hanamaru.Context) error {
		input, err := ctx.GetArgIndex(0)
		if err != nil {
			return err
		}
		image, err := latex.GeneratePNGFromLatex(input)
		if err != nil {
			return err
		}
		ctx.ReplyFile("latex.png", image)
		return nil
	},
}
