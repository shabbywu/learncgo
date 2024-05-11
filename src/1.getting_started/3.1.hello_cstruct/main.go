package main

// #include "lib.h"
import "C"
import (
    "fmt"
    "reflect"
)


func main() {
    {
        fmt.Println("[*] new struct _Foo instance, and show all fields in Foo type reflect")
        foo := C.struct__Foo{
            foo: 1,
            bar: 2.2,
            baz: 3.3,

            _type: 1,
            __type: 2,
            ___type: 3,
        }

        t := reflect.TypeOf(foo)
        for i := 0; i < t.NumField(); i++ {
            // Get FieldName and FieldType
            field := t.Field(i)
            name := field.Name
            fieldType := field.Type
            // Get Field Value
            value := reflect.ValueOf(foo).FieldByName(name)
            fmt.Printf("%s (%s) = %v\n", name, fieldType, value)
        }
    }

    fmt.Println("==========================")

    {
        fmt.Println("[*] new struct _Foo instance via typedef alias")
        foo := C.Foo{
            foo: 1,
            bar: 2.2,
            baz: 3.3,

            _type: 1,
            __type: 2,
            ___type: 3,
        }
        fmt.Println("foo.foo:", foo.foo)
        fmt.Println("foo.bar:", foo.bar)
        fmt.Println("foo.baz:", foo.baz)
        fmt.Println("foo.func:", foo._func)
        fmt.Println("foo._func:", foo.__func)

        fmt.Println("foo.type:", foo._type)
        fmt.Println("foo._type:", foo.__type)
        fmt.Println("foo.__type:", foo.___type)
    }

    fmt.Println("==========================")

    {
        fmt.Println("[*] new struct _Foo instance via c api")
        foo := C.Foo_init(1, 2.2, 3.3)
        fmt.Println("foo.foo:", foo.foo)
        fmt.Println("foo.bar:", foo.bar)
        fmt.Println("foo.baz:", foo.baz)
        fmt.Println("foo.func:", foo._func)
        fmt.Println("foo._func:", foo.__func)

        fmt.Println("foo.type:", foo._type)
        fmt.Println("foo._type:", foo.__type)
        fmt.Println("foo.__type:", foo.___type)
    }

    fmt.Println("==========================")

    {
        fmt.Println("[*] new struct _Foo instance(pointer) via c api")
        foop := C.Foo_new(1, 2.2, 3.3)
        fmt.Println("foop.foo:", foop.foo)
        fmt.Println("foop.bar:", foop.bar)
        fmt.Println("foop.baz:", foop.baz)
        fmt.Println("foop.func:", foop._func)
        fmt.Println("foop._func:", foop.__func)

        fmt.Println("foop.type:", foop._type)
        fmt.Println("foop._type:", foop.__type)
        fmt.Println("foop.__type:", foop.___type)
        C.Foo_free(foop);

        // Undefined behavior
        // fmt.Println(foop.foo, foop.bar, foop.baz)
        // Output: 1980710944 4.7292e-41 2e-323
    }
}
