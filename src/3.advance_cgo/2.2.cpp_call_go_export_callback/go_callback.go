package main

// #include <stdint.h>
import "C"

//export GoDouble
func GoDouble(value C.int) C.int {
	return value * 2
}

func main() {}
