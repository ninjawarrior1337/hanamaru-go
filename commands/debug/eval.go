package debug

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
	"strings"
	"time"
)

var Eval = &framework.Command{
	Name:               "eval",
	PermissionRequired: discordgo.PermissionAdministrator,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		L := lua.NewState(lua.Options{
			SkipOpenLibs: true,
		})

		lua.OpenBase(L)
		lua.OpenMath(L)
		lua.OpenString(L)
		lua.OpenTable(L)

		goCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		L.SetContext(goCtx)

		defer L.Close()
		L.SetGlobal("print", L.NewFunction(func(state *lua.LState) int {
			top := L.GetTop()
			for i := 1; i <= top; i++ {
				ctx.Reply(L.ToStringMeta(L.Get(i)).String())
			}
			return 0
		}))
		L.SetGlobal("this", luar.New(L, ctx.Message))
		var luaCtx *framework.Context
		if ctx.Author.ID != ctx.Hanamaru.GetOwnerID() {
			luaCtx = &framework.Context{
				Hanamaru:      nil,
				Command:       ctx.Command,
				MessageCreate: ctx.MessageCreate,
			}
		} else {
			luaCtx = ctx
		}
		L.SetGlobal("ctx", luar.New(L, luaCtx))
		luaCode := strings.Replace(ctx.TakeRest(), "lua", "", 1)
		if err := L.DoString(luaCode); err != nil {
			return err
		}
		return nil
	},
	Setup: nil,
}
