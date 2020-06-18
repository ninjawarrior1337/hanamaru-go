package fun

import (
	"github.com/markbates/pkger"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"image"
	"math/rand"
	"os"
)

const PadoruDir = "/assets/imgs/padoru"

var padoruList = make([]string, 0)

var Padoru = &framework.Command{
	Name:               "padoru",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		imgFile, _ := pkger.Open(PadoruDir + "/" + padoruList[rand.Intn(len(padoruList))])
		defer imgFile.Close()
		img, _, _ := image.Decode(imgFile)
		ctx.ReplyJPGImg(img, imgFile.Name())
		return nil
	},
	Setup: func() error {
		err := pkger.Walk(PadoruDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			padoruList = append(padoruList, info.Name())
			return nil
		})
		return err
	},
}
