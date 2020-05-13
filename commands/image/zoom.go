package image

import (
	"fmt"
	"hanamaru/hanamaru"
	"image/gif"
	"net/http"
)

var Zoom = &hanamaru.Command{
	Name:               "zoom",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *hanamaru.Context) error {
		msg, err := ctx.GetPreviousMessage()
		if err != nil {
			return err
		}
		if len(msg.Embeds) <= 0 {
			return fmt.Errorf("this command can only be used with embeded images")
		}
		imgUrl := msg.Embeds[0].URL
		resp, err := http.Get(imgUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		g, err := gif.DecodeAll(resp.Body)
		if err != nil {
			return err
		}

		for i, d := range g.Delay {
			g.Delay[i] = d / 2
		}

		_, err = ctx.ReplyGIFImg(g, "zoom")
		if err != nil {
			return err
		}

		return nil
	},
}
