package main

import (
	"fmt"

	"github.com/shabbywu/learncgo/src/3.advance_cgo/containers"
)

func main() {
	values, err := containers.NewArray()
	if err != nil {
		panic(err)
	}
	defer values.Close()
	for index, value := range []int{2, 3, 5, 7} {
		values.Set(index, value)
	}
	if got := values.Values(); fmt.Sprint(got) != "[2 3 5 7]" {
		panic(fmt.Sprintf("unexpected array: %v", got))
	}
	fmt.Println("OK 3.advance_cgo_1.1.cpp_std_array")
}
