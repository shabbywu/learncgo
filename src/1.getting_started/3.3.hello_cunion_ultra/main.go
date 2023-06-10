package main

// #include "lib.h"
import "C"
import "fmt"
import "unsafe"


func set_union[T any, VT any](o *T, v VT){
	ptr := (*T)(unsafe.Pointer(o))
	*ptr = *(*T)(unsafe.Pointer(&v))
}


func read_union[T any, OT any](o OT) T {
	ptr := (*T)(unsafe.Pointer(&o))
	return *ptr
}


func main() {
	foo := C.Foo{}
	set_union(&foo.u, C.int(1))
	fmt.Println("read as int=", read_union[C.int](foo.u))
	fmt.Println("read as float=", read_union[C.float](foo.u))
	fmt.Println("read as double=", read_union[C.double](foo.u))

	set_union(&foo.u, C.float(2.2))
	fmt.Println("read as int=", read_union[C.int](foo.u))
	fmt.Println("read as float=", read_union[C.float](foo.u))
	fmt.Println("read as double=", read_union[C.double](foo.u))

	set_union(&foo.u, C.double(3.3))
	fmt.Println("read as int=", read_union[C.int](foo.u))
	fmt.Println("read as float=", read_union[C.float](foo.u))
	fmt.Println("read as double=", read_union[C.double](foo.u))
}