package debug

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"

	"github.com/dop251/goja"
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

var testEthersScript = `
const provider = new ethers.providers.CloudflareProvider();
const g = async () => {
	let b = await provider.getBlock(0)
	console.log(b)
}
g()
`

func newRandSource() goja.RandSource {
	var seed int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		panic(fmt.Errorf("Could not read random bytes: %v", err))
	}
	return rand.New(rand.NewSource(seed)).Float64
}

func TestEthers(t *testing.T) {
	vm := goja.New()
	setupConsole(vm)
	setupEthers(vm)
	vm.SetRandSource(newRandSource())
	transpile, _ := setupBabel(vm)
	transpiled, err := transpile(testEthersScript)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(transpiled)
	v, err := vm.RunString(transpiled)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
