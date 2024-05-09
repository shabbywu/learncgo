package main

// #include "lib.h"
import "C"
import "fmt"


func main() {
    foo_int := C.Foo_init_foo(1)
    fmt.Println("int 1=", C.read_u_foo(foo_int.u))
    fmt.Println("int 1 as flaot=", C.read_u_bar(foo_int.u))
    fmt.Println("int 1 as double=", C.read_u_baz(foo_int.u))

    foo_float := C.Foo_init_bar(2.2)
    fmt.Println("float 2.2 as int=", C.read_u_foo(foo_float.u))
    fmt.Println("float 2.2=", C.read_u_bar(foo_float.u))
    fmt.Println("float 2.2 as double=", C.read_u_baz(foo_float.u))

    foo_double := C.Foo_init_baz(3.3)
    fmt.Println("double 3.3 as int=", C.read_u_foo(foo_double.u))
    fmt.Println("double 3.3 as flaot=", C.read_u_bar(foo_double.u))
    fmt.Println("double 3.3=", C.read_u_baz(foo_double.u))

    // 当然，也可以直接构造 Foo
    foo := C.Foo{[8]byte{1,2,3,4,5,6,7,8}}
    fmt.Println("foo as int=", C.read_u_foo(foo.u))
    fmt.Println("foo as flaot=", C.read_u_bar(foo.u))
    fmt.Println("foo as double=", C.read_u_baz(foo.u))
}
