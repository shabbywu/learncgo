package main

// #include "lib.h"
import "C"
import "fmt"

func main() {
    fmt.Println("RED=", C.RED)
    fmt.Println("GREEN=", C.GREEN)
    fmt.Println("BLUE=", C.BLUE)
    fmt.Println("ALPHA=", C.ALPHA)
    var o C.enum_RGBA = C.BLUE
    fmt.Println("o=", o)
}
