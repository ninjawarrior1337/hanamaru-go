package fun

import (
	"image"
	"os"

	"github.com/ninjawarrior1337/hanamaru-go/util"

	"github.com/markbates/pkger"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

const padoruDir = "/assets/imgs/padoru"

var padoruList = make([]string, 0)
var prevPick = 0

var Padoru = &framework.Command{
	Name:               "padoru",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		imgFile, _ := pkger.Open(padoruDir + "/" + padoruList[util.IntnNoDup(len(padoruList), &prevPick)])
		defer imgFile.Close()
		img, _, _ := image.Decode(imgFile)
		ctx.ReplyJPGImg(img, imgFile.Name())
		return nil
	},
	Setup: func() error {
		err := pkger.Walk(padoruDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			padoruList = append(padoruList, info.Name())
			return nil
		})
		return err
	},
}
