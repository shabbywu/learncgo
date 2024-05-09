package main

// #include "lib.h"
import "C"
import "fmt"

func main() {
    // 创建 MyClass 对象
    myObject := C.MyClass_new(42)

    // cgo 不能直接访问 c++ 对象
    // fmt.Println(myObject.m_value)
    fmt.Println("value=", C.MyClass_GetValue(myObject))

    // 调用 MyClass 的方法
    C.MyClass_SetValue(myObject, 100)
    fmt.Println("value=", C.MyClass_GetValue(myObject))

    // 销毁 MyClass 对象
    C.MyClass_delete(myObject)
    // 未定义行为
    // fmt.Println("value=", C.MyClass_GetValue(myObject))
    // Output: 1980710944
}
