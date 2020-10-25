package fun

import (
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
)

var FlagEnc = &framework.Command{
	Name:               "flagenc",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Generates a flag...thats it <https://en.wikipedia.org/wiki/Illegal_number>",
	Exec: func(ctx *framework.Context) error {
		data := ctx.TakeRest()
		flag := util.EncodeFlag([]byte(data))
		ctx.ReplyPNGImg(flag, "flag")
		return nil
	},
	Setup: nil,
}

var FlagDec = &framework.Command{
	Name:               "flagdec",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Extracts the data from a flag...thats it <https://en.wikipedia.org/wiki/Illegal_number>",
	Exec: func(ctx *framework.Context) error {
		flag, err := ctx.GetImage(0)
		if err != nil {
			return err
		}
		data, err := util.DecodeFlag(flag)
		if err != nil {
			return err
		}
		if string(data) == "" {
			ctx.Reply("This flag contains no data.")
			return nil
		}
		ctx.Reply(string(data))
		return nil
	},
	Setup: nil,
}
