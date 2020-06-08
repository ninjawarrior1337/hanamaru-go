package image

import (
	"github.com/disintegration/imaging"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util/latex"
)

var Latex = &framework.Command{
	Name:               "latex",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Exec: func(ctx *framework.Context) error {
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
