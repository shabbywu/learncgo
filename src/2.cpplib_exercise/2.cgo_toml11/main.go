package main

// #include "toml11_lib.h"
import "C"
import "fmt"


const content string = `
title = "an example toml file"
nums  = [3, 1, 4, 1, 5]
`

func main() {
    c_content := C.CString(content)
    h := C.toml11_parse_content(c_content)

    c_title := C.CString("title")
    fmt.Printf("title: %s\n", C.GoString(C.toml11_find_string(h, c_title)))

    C.toml11_free(h);
    // 未定义行为, 内存已释放, 继续使用指针会 crash
    // fmt.Printf("title: %s\n", C.GoString(C.toml11_find_string(h, c_title)))
}
