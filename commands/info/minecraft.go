package info

import (
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
)

var Minecraft = &framework.Command{
	Name:               "mc",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		mcName, err := ctx.GetArgIndex(0)
		rt := ctx.GetArgIndexDefault(1, util.Body)

		if err != nil {
			return err
		}
		p, err := util.LookupMinecraft(mcName)
		if err != nil {
			return fmt.Errorf("player not found: %v", mcName)
		}
		img, err := util.GetMinecraftSkin(p, util.MinecraftRenderType(rt))
		if err != nil {
			return err
		}
		ctx.ReplyPNGImg(img, mcName)
		return nil
	},
}
