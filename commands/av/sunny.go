package av

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/ninjawarrior1337/hanamaru-go/util"
	"github.com/ninjawarrior1337/hanamaru-go/util/av"
)

var Sunny = &framework.Command{
	Name: "sunny",
	Help: "Generates always sunny in philadelphia title cards",
	Exec: func(ctx *framework.Context) error {
		arg := ctx.TakeRest()
		arg = arg[1:]
		board := gg.NewContext(1920, 1080)
		board.SetColor(color.Black)
		board.Fill()
		board.SetColor(color.White)
		board.SetFontFace(util.GetFontByName("Textile", 128))
		board.DrawStringWrapped(fmt.Sprintf(`"%v"`, arg), float64(board.Width()/2), float64(board.Height()/2), 0.5, 0.5, float64(board.Width())/1.75, 1, gg.AlignCenter)

		ctx.ReplyJPGImg(board.Image(), "phily")

		return nil
	},
}

var SunnyAudio = &framework.Command{
	Name: "asunny",
	Help: "Generates always sunny in philadelphia title cards with audio (experimental)",
	Exec: func(ctx *framework.Context) error {
		arg := ctx.TakeRest()
		arg = arg[1:]
		board := gg.NewContext(1920, 1080)
		board.SetColor(color.Black)
		board.Fill()
		board.SetColor(color.White)
		board.SetFontFace(util.GetFontByName("Textile", 128))
		board.DrawStringWrapped(fmt.Sprintf(`"%v"`, arg), float64(board.Width()/2), float64(board.Height()/2), 0.5, 0.5, float64(board.Width())/1.75, 1, gg.AlignCenter)

		videoBuffer := av.AddAudioToImage(board.Image(), "sunny.ogg")
		ctx.ReplyFile("sunny.webm", videoBuffer)
		return nil
	},
}

var SunnyVideo = &framework.Command{
	Name: "vsunny",
	Help: "Generates always sunny in philadelphia title sequences (experimental)",
	Exec: func(ctx *framework.Context) error {
		arg := ctx.TakeRest()
		arg = arg[1:]
		board := gg.NewContext(1920, 1080)
		board.SetColor(color.Black)
		board.Fill()
		board.SetColor(color.White)
		board.SetFontFace(util.GetFontByName("Textile", 128))
		board.DrawStringWrapped(fmt.Sprintf(`"%v"`, arg), float64(board.Width()/2), float64(board.Height()/2), 0.5, 0.5, float64(board.Width())/1.75, 1, gg.AlignCenter)

		videoBuffer := av.OverlayImage(board.Image(), "sunny.webm", 0, 90)
		ctx.ReplyFile("sunny.mp4", videoBuffer)
		return nil
	},
}
