package main

// #cgo LDFLAGS: -lhandle -lstdc++
// #include "handle.h"
import "C"

import (
	"fmt"
	"runtime/cgo"
)

var callbackResult string

//export goHandleCallback
func goHandleCallback(value C.uintptr_t) {
	callbackResult = cgo.Handle(value).Value().(string)
}

func main() {
	value := cgo.NewHandle("kept by a uintptr_t, not a Go pointer")
	native := C.handle_new(C.uintptr_t(value))
	if native == nil {
		value.Delete()
		panic("native allocation failed")
	}
	C.handle_invoke(native)
	C.handle_free(native) // Native code no longer retains the handle value.
	value.Delete()
	if callbackResult != "kept by a uintptr_t, not a Go pointer" {
		panic("unexpected handle callback value")
	}
	fmt.Println("OK 3.advance_cgo_2.3.cgo_handle")
}
