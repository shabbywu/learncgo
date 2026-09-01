package main

import (
	"fmt"

	"github.com/shabbywu/learncgo/src/3.advance_cgo/containers"
)

func main() {
	values, err := containers.NewSet()
	if err != nil {
		panic(err)
	}
	defer values.Close()
	if !values.Insert(5) || !values.Insert(2) || values.Insert(5) || !values.Contains(2) {
		panic("unexpected set behavior")
	}
	if got := values.Values(); fmt.Sprint(got) != "[2 5]" {
		panic(fmt.Sprintf("unexpected set: %v", got))
	}
	fmt.Println("OK 3.advance_cgo_1.3.cpp_std_set")
}
