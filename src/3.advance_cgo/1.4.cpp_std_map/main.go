package main

import (
	"fmt"

	"github.com/shabbywu/learncgo/src/3.advance_cgo/containers"
)

func main() {
	values, err := containers.NewMap()
	if err != nil {
		panic(err)
	}
	defer values.Close()
	values.Set("a\\x00b", 8)
	values.Set("answer", 42)
	if value, found := values.Get("a\\x00b"); !found || value != 8 {
		panic("embedded NUL key was not preserved")
	}
	if value, found := values.Get("answer"); !found || value != 42 {
		panic("unexpected map value")
	}
	fmt.Println("OK 3.advance_cgo_1.4.cpp_std_map")
}
