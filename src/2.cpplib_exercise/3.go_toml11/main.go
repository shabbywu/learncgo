package main

// #include "toml11_lib.h"
// #include <stdlib.h>
import "C"
import "fmt"
import "learncgo/go_toml11/toml11"


const content string = `
title = "an example toml file"
nums  = [3, 1, 4, 1, 5]
`

func main() {
    h := toml11.Parse(content)
    fmt.Printf("title: %s\n", toml11.FindString(h, "title"))
    toml11.Free(h)
}
