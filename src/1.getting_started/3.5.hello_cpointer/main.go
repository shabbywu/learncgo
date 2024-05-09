package main

// #include "lib.h"
import "C"
import "fmt"

func main() {
    foo := C.Foo_init(1, 2.2, 3.3)
    fmt.Println("object", foo.foo, foo.bar, foo.baz)

    foop := C.Foo_new(1, 2.2, 3.3)
    fmt.Println("pointer", foop.foo, foop.bar, foop.baz)

    C.Foo_free(foop);
    // 未定义行为
    // fmt.Println(foop.foo, foop.bar, foop.baz)
    // Output: 1980710944 4.7292e-41 2e-323
}
