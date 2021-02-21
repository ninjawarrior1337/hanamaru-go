package fun

import (
	"embed"
	"errors"
	"io"
	"strconv"

	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

//go:embed assets/spoonful/*
var spoonfulFS embed.FS

func getSpoonfulImage(fileName string) io.Reader {
	f, _ := spoonfulFS.Open("assets/spoonful/" + fileName)
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
