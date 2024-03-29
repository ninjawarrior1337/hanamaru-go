package debug

/*
#cgo LDFLAGS: -L../../lib/target/release -lhanamaru_lib -lm
#include <stdlib.h>
#include <stdint.h>
extern void goSendMessage(uintptr_t handle, char* msg);
extern char* goReferencedMessage(uintptr_t handle);

char* run_uiua(uintptr_t ctx, char* uiua);
void drop_string(char* ptr);
*/
import "C"
import (
	"fmt"
	"runtime/cgo"
	"unsafe"

	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

//export goSendMessage
func goSendMessage(handle C.uintptr_t, message *C.char) {
	h := cgo.Handle(handle)

	ctx := h.Value().(*framework.Context)
	ctx.Reply(fmt.Sprintf("```\n%v\n```", C.GoString(message)))
	C.drop_string(message)
}

//export goReferencedMessage
func goReferencedMessage(handle C.uintptr_t) *C.char {
	h := cgo.Handle(handle)

	ctx := h.Value().(*framework.Context)
	if ctx.Message.ReferencedMessage != nil {
		return C.CString(ctx.Message.ReferencedMessage.Content)
	} else {
		return nil
	}
}

//export goAssignRole
func goAssignRole(handle C.uintptr_t, targetId_c *C.char, roleId_c *C.char) {
	h := cgo.Handle(handle)
	ctx := h.Value().(*framework.Context)

	targetId := C.GoString(targetId_c)
	defer C.drop_string(targetId_c)
	roleId := C.GoString(roleId_c)
	defer C.drop_string(roleId_c)

	err := ctx.Hanamaru.GuildMemberRoleAdd(ctx.GuildID, targetId, roleId)
	if err != nil {
		ctx.Reply(fmt.Sprintf("ERROR: failed to set role %v", err))
	}
}

var Uiua = &framework.Command{
	Name:               "uiua",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		input := ctx.TakeRest()
		h := cgo.NewHandle(ctx)
		defer h.Delete()

		input_c := C.CString(input)
		defer C.free(unsafe.Pointer(input_c))

		res_c := C.run_uiua(C.uintptr_t(h), input_c)
		defer C.drop_string(res_c)

		res := C.GoString(res_c)
		_, err := ctx.Reply(fmt.Sprintf("```\n%v\n```", res))
		if err != nil {
			return err
		}

		return nil
	},
}
