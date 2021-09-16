package debug

import (
	_ "embed"
	"fmt"
	"github.com/dop251/goja"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"time"
)

//go:embed babel.min.js
var babel string

var vm = goja.New()
var transpile func(string) (string, error)

func setupConsole(vm *goja.Runtime) {
	logFunc := func(goja.FunctionCall) goja.Value { return nil }
	vm.Set("console", map[string]func(goja.FunctionCall) goja.Value{
		"log":   logFunc,
		"error": logFunc,
		"warn":  logFunc,
	})
}

func init() {
	setupConsole(vm)
	_, err := vm.RunScript("babel", babel)
	if err != nil {
		panic(err)
	}
	transform, _ := goja.AssertFunction(vm.Get("Babel").ToObject(vm).Get("transform"))
	configObject := vm.NewObject()
	configObject.Set("presets", []string{"env"})
	configObject.Set("sourceType", "unambiguous")
	transpile = func(s string) (string, error) {
		transformResult, _ := transform(goja.Undefined(), vm.ToValue(s), configObject)
		if transformResult != nil {
			return transformResult.ToObject(vm).Get("code").String(), nil
		} else {
			return "", fmt.Errorf("failed to compile")
		}
	}
}

var Eval = &framework.Command{
	Name:               "eval",
	PermissionRequired: 0,
	OwnerOnly:          true,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		t := time.AfterFunc(2 * time.Second, func() {
			vm.Interrupt("Timeout")
		})
		defer t.Stop()
		defer vm.ClearInterrupt()
		vm.Set("print", func(v... interface{}) {
			ctx.Reply(fmt.Sprintf("%v", v))
		})
		vm.Set("ctx", ctx)

		transpiledJS, err := transpile(ctx.TakeRest())
		if err != nil {
			return err
		}
		fmt.Println("Running: "+transpiledJS)
		v, err := vm.RunString(transpiledJS)
		if err != nil {
			return err
		}
		if len(v.String()) > 0 {
			ctx.Reply(v.String())
		}
		return nil
	},
	Setup: nil,
}
