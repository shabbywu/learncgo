package main

// #cgo LDFLAGS: -lcallback -lstdc++
// #include "callback.h"
//
// // cgo cannot invoke a C function pointer directly. This C shim supplies the
// // C++ function as a pointer to the native bridge, which performs the call.
// static int call_cpp_increment(int value) {
//   return invoke_int_callback(cpp_increment, value);
// }
import "C"

import "fmt"

func main() {
	if got := int(C.call_cpp_increment(41)); got != 42 {
		panic(fmt.Sprintf("unexpected callback result: %d", got))
	}
	fmt.Println("OK 3.advance_cgo_2.1.go_call_cpp_function_pointer")
}
