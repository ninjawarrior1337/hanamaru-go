package fun

import (
	"embed"
	"image"

	"github.com/ninjawarrior1337/hanamaru-go/util"

	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

//go:embed assets/padoru/*
var padoruFS embed.FS

var padoruList = make([]string, 0)
var prevPick = 0

var Padoru = &framework.Command{
	Name:               "padoru",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		imgFile, _ := padoruFS.Open("assets/padoru" + "/" + padoruList[util.IntnNoDup(len(padoruList), &prevPick)])
		defer imgFile.Close()
		img, _, _ := image.Decode(imgFile)
		fileInfo, _ := imgFile.Stat()
		ctx.ReplyJPGImg(img, fileInfo.Name())
		return nil
	},
	Setup: func() error {
		dir, err := padoruFS.ReadDir("assets/padoru")
		for _, f := range dir {
			padoruList = append(padoruList, f.Name())
		}
		return err
	},
}
