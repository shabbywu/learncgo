package main

// #include "lib.h"
import "C"
import "unsafe"

func main() {
    cs := C.CString("Hello from stdio")
    C.CPrintf(cs)
    C.free(unsafe.Pointer(cs))
}
