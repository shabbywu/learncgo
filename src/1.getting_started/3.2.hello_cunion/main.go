package main

// #include "lib.h"
import "C"
import (
    "fmt"
    "reflect"
)


func main() {

    {
        fmt.Println("[*] new union _u instance, and show the go type of union _u")
        u := C.union__u{1,2,3,4,5,6,7,8}

        t := reflect.TypeOf(u)
        fmt.Printf("type union _u in go is %+v\n", t)
    }

    fmt.Println("==========================")

    {
        fmt.Println("[*] new union _u instance via typedef alias, and access it via c api")
        u := C.U{1,2,3,4,5,6,7,8}
        fmt.Println("u as int=", C.read_u_foo(u))
        fmt.Println("u as flaot=", C.read_u_bar(u))
        fmt.Println("u as double=", C.read_u_baz(u))
    }

    {
        fmt.Println("[*] new union _u via go type, and access it via c api")
        u := [8]byte{1,2,3,4,5,6,7,8}
        fmt.Println("u as int=", C.read_u_foo(u))
        fmt.Println("u as flaot=", C.read_u_bar(u))
        fmt.Println("u as double=", C.read_u_baz(u))
    }

    fmt.Println("==========================")

    {
        fmt.Println("[*] new union u instance in struct Foo, and access it via c api")
        foo := C.Foo{[8]byte{1,2,3,4,5,6,7,8}}
        fmt.Println("foo.u as int=", C.read_u_foo(foo.u))
        fmt.Println("foo.u as flaot=", C.read_u_bar(foo.u))
        fmt.Println("foo.u as double=", C.read_u_baz(foo.u))
    }

    fmt.Println("==========================")

    {
        fmt.Println("[*] new union u instance in struct Foo via c api, and access it via c api")
        fmt.Println("  [*] case Foo.u = 1 (int)")
        foo_int := C.Foo_init_foo(1)
        fmt.Println("int 1=", C.read_u_foo(foo_int.u))
        fmt.Println("int 1 as flaot=", C.read_u_bar(foo_int.u))
        fmt.Println("int 1 as double=", C.read_u_baz(foo_int.u))
        
        fmt.Println("  [*] case Foo.u = 2.2 (float)")
        foo_float := C.Foo_init_bar(2.2)
        fmt.Println("float 2.2 as int=", C.read_u_foo(foo_float.u))
        fmt.Println("float 2.2=", C.read_u_bar(foo_float.u))
        fmt.Println("float 2.2 as double=", C.read_u_baz(foo_float.u))
    
        fmt.Println("  [*] case Foo.u = 3.3 (double)")
        foo_double := C.Foo_init_baz(3.3)
        fmt.Println("double 3.3 as int=", C.read_u_foo(foo_double.u))
        fmt.Println("double 3.3 as flaot=", C.read_u_bar(foo_double.u))
        fmt.Println("double 3.3=", C.read_u_baz(foo_double.u))
    }
}
