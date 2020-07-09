package fun

import (
	"errors"
	"github.com/markbates/pkger"
	"github.com/markbates/pkger/pkging"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"strconv"
)

func getSpoonfulImage(fileName string) pkging.File {
	pkger.Include("/assets/imgs/spoonful")
	f, _ := pkger.Open("/assets/imgs/spoonful/" + fileName)
	return f
}

var Spoonful = &framework.Command{
	Name:               "spoonful",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Prints the spoonful image except with minecraft hoe, similar to patrick command",
	Exec: func(ctx *framework.Context) error {
		length, _ := strconv.Atoi(ctx.GetArgIndexDefault(0, "0"))
		if length <= 0 {
			ctx.ReplyFile("0.jpg", getSpoonfulImage("0.jpg"))
			ctx.ReplyFile("2.jpg", getSpoonfulImage("2.jpg"))
		} else if length <= 3 {
			ctx.ReplyFile("0.jpg", getSpoonfulImage("0.jpg"))
			for i := 0; i < length; i++ {
				ctx.ReplyFile("1.jpg", getSpoonfulImage("1.jpg"))
			}
			ctx.ReplyFile("2.jpg", getSpoonfulImage("2.jpg"))
		} else {
			return errors.New("length too long")
		}
		return nil
	},
}
