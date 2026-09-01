package main

// #include "lib.h"
import "C"
import "fmt"
import "unsafe"

func setInt(dst unsafe.Pointer, value C.int)       { *(*C.int)(dst) = value }
func setFloat(dst unsafe.Pointer, value C.float)   { *(*C.float)(dst) = value }
func setDouble(dst unsafe.Pointer, value C.double) { *(*C.double)(dst) = value }

func readInt(src unsafe.Pointer) C.int       { return *(*C.int)(src) }
func readFloat(src unsafe.Pointer) C.float   { return *(*C.float)(src) }
func readDouble(src unsafe.Pointer) C.double { return *(*C.double)(src) }

func main() {
	foo := C.Foo{}
	setInt(unsafe.Pointer(&foo.u), C.int(1))
	fmt.Println("read as int=", readInt(unsafe.Pointer(&foo.u)))
	fmt.Println("read as float=", readFloat(unsafe.Pointer(&foo.u)))
	fmt.Println("read as double=", readDouble(unsafe.Pointer(&foo.u)))

	setFloat(unsafe.Pointer(&foo.u), C.float(2.2))
	fmt.Println("read as int=", readInt(unsafe.Pointer(&foo.u)))
	fmt.Println("read as float=", readFloat(unsafe.Pointer(&foo.u)))
	fmt.Println("read as double=", readDouble(unsafe.Pointer(&foo.u)))

	setDouble(unsafe.Pointer(&foo.u), C.double(3.3))
	fmt.Println("read as int=", readInt(unsafe.Pointer(&foo.u)))
	fmt.Println("read as float=", readFloat(unsafe.Pointer(&foo.u)))
	fmt.Println("read as double=", readDouble(unsafe.Pointer(&foo.u)))
}
