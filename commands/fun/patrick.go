package fun

import (
	"github.com/markbates/pkger"
	"github.com/markbates/pkger/pkging"
	"hanamaru/hanamaru"
	"strconv"
)

func getPatrickImage(fileName string) pkging.File {
	pkger.Include("/assets/imgs/patrick")
	f, _ := pkger.Open("/assets/imgs/patrick/" + fileName)
	return f
}

var Patrick = &hanamaru.Command{
	Name:               "patrick",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *hanamaru.Context) error {
		neck, _ := strconv.Atoi(ctx.GetArgIndexDefault(0, "0"))
		if neck < 0 {
			ctx.ReplyFile("0r.jpg", getPatrickImage("0r.jpg"))
			for i := 0; i > neck; i-- {
				ctx.ReplyFile("1r.jpg", getPatrickImage("1r.jpg"))
			}
			ctx.ReplyFile("2r.jpg", getPatrickImage("2r.jpg"))
		} else if neck > 0 {
			ctx.ReplyFile("0.jpg", getPatrickImage("0.jpg"))
			for i := 0; i < neck; i++ {
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
