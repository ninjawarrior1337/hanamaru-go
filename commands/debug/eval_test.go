package debug

import (
	"fmt"
	"github.com/dop251/goja"
	"testing"
)

func TestVM(t *testing.T) {
	vm := goja.New()
	setupConsole(vm)
	vm.RunScript("babel", babel)
	transform, _ := goja.AssertFunction(vm.Get("Babel").ToObject(vm).Get("transform"))
	configObject := vm.NewObject()
	configObject.Set("presets", []string{"env"})
	configObject.Set("sourceType", "unambiguous")
	transformResult, _ := transform(goja.Undefined(), vm.ToValue(`const x = 0;`), configObject)
	if transformResult != nil {
		fmt.Println(transformResult.ToObject(vm).Get("code"))
	}
}
