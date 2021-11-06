package debug

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dop251/goja"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

//go:embed babel.min.js
var babel string

var ethers string

var vm = goja.New()

type TranspileFunc func(string) (string, error)

var transpile TranspileFunc

func setupConsole(vm *goja.Runtime) error {
	logFunc := func(f goja.FunctionCall) goja.Value {
		fmt.Println(f.Arguments)
		return goja.Undefined()
	}
	return vm.Set("console", map[string]func(goja.FunctionCall) goja.Value{
		"log":   logFunc,
		"error": logFunc,
		"warn":  logFunc,
	})
}

func setupEthers(vm *goja.Runtime) error {
	resp, err := http.Get("https://cdn.ethers.io/lib/ethers-5.2.umd.min.js")
	if err != nil {
		return err
	}
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	ethers = string(respBodyBytes)
	_, err = vm.RunScript("ethers", ethers)
	if err != nil {
		return err
	}
	return nil
}

func setupBabel(vm *goja.Runtime) (TranspileFunc, error) {
	_, err := vm.RunScript("babel", babel)
	if err != nil {
		return nil, err
	}
	transform, _ := goja.AssertFunction(vm.Get("Babel").ToObject(vm).Get("transform"))
	configObject := vm.NewObject()
	configObject.Set("presets", []string{"env"})
	configObject.Set("sourceType", "unambiguous")
	return func(s string) (string, error) {
		transformResult, _ := transform(goja.Undefined(), vm.ToValue(s), configObject)
		if transformResult != nil {
			return transformResult.ToObject(vm).Get("code").String(), nil
		} else {
			return "", fmt.Errorf("failed to compile")
		}
	}, nil
}

func setup() error {
	err := setupConsole(vm)
	if err != nil {
		return err
	}
	err = setupEthers(vm)
	if err != nil {
		return err
	}
	transpile, err = setupBabel(vm)
	if err != nil {
		return err
	}

	return nil
}

var Eval = &framework.Command{
	Name:               "eval",
	PermissionRequired: 0,
	OwnerOnly:          true,
	Help:               "",
	Exec: func(ctx *framework.Context) error {
		t := time.AfterFunc(2*time.Second, func() {
			vm.Interrupt("Timeout")
		})
		defer t.Stop()
		defer vm.ClearInterrupt()
		vm.Set("print", func(v ...interface{}) {
			ctx.Reply(fmt.Sprintf("%v", v))
		})
		vm.Set("ctx", ctx)

		transpiledJS, err := transpile(ctx.TakeRest())
		if err != nil {
			return err
		}
		fmt.Println("Running: " + transpiledJS)
		v, err := vm.RunString(transpiledJS)
		if err != nil {
			return err
		}
		if len(v.String()) > 0 {
			ctx.Reply(v.String())
		}
		return nil
	},
	Setup: setup,
}
