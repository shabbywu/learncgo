package toml11

// #include "../toml11_lib.h"
// #include <stdlib.h>
import "C"
import "unsafe"


type Toml11Value struct {
    Pointer unsafe.Pointer
}


func Parse(content string) Toml11Value {
    c_content := C.CString(content)
    h := C.toml11_parse_content(c_content)
    C.free(unsafe.Pointer(c_content))
    return Toml11Value{h}
}


func Free(handler Toml11Value) {
    C.toml11_free(handler.Pointer)
}

func FindString(handler Toml11Value, key string) string {
    c_key := C.CString(key)
    v := C.GoString(C.toml11_find_string(handler.Pointer, c_key))
    C.free(unsafe.Pointer(c_key))
    return v
}
