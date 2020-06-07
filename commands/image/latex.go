package image

import (
	"github.com/disintegration/imaging"
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
		image, err := latex.GenerateLatexImage(input)
		neg := imaging.Invert(image)
		if err != nil {
			return err
		}
		ctx.ReplyPNGImg(neg, "latex")
		return nil
	},
}
