package fun

import (
	"embed"
	"io"
	"strconv"

	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

//go:embed assets/patrick/*
var patrickFS embed.FS

func getPatrickImage(fileName string) io.Reader {
	f, _ := patrickFS.Open("assets/patrick/" + fileName)
	return f
}

var Patrick = &framework.Command{
	Name:               "patrick",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		length, _ := strconv.Atoi(ctx.GetArgIndexDefault(0, "0"))
		if length < 0 {
			ctx.ReplyFile("0r.jpg", getPatrickImage("0r.jpg"))
			for i := 0; i > length; i-- {
				ctx.ReplyFile("1r.jpg", getPatrickImage("1r.jpg"))
			}
			ctx.ReplyFile("2r.jpg", getPatrickImage("2r.jpg"))
		} else if length > 0 {
			ctx.ReplyFile("0.jpg", getPatrickImage("0.jpg"))
			for i := 0; i < length; i++ {
				ctx.ReplyFile("1.jpg", getPatrickImage("1.jpg"))
			}
			ctx.ReplyFile("2.jpg", getPatrickImage("2.jpg"))
		} else {
			ctx.ReplyFile("0.jpg", getPatrickImage("0.jpg"))
			ctx.ReplyFile("2.jpg", getPatrickImage("2.jpg"))
		}
		return nil
	},
}
