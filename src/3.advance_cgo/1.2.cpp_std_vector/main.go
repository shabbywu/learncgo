package main

import (
	"fmt"

	"github.com/shabbywu/learncgo/src/3.advance_cgo/containers"
)

func main() {
	values, err := containers.NewVector()
	if err != nil {
		panic(err)
	}
	defer values.Close()
	for _, value := range []int{3, 1, 4} {
		values.Push(value)
	}
	if got := values.Values(); fmt.Sprint(got) != "[3 1 4]" {
		panic(fmt.Sprintf("unexpected vector: %v", got))
	}
	fmt.Println("OK 3.advance_cgo_1.2.cpp_std_vector")
}
