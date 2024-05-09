package main

// #include "lib.h"
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
    cs := C.CString("Hello from iostream")
    cpps := C.CppString(cs, 20)
    C.CppPrint(cpps)
    C.CppStringFree(cpps)
    C.free(unsafe.Pointer(cs))
}
