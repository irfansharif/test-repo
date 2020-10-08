package main

import "fmt"

// #cgo CPPFLAGS: -Ic-deps/libroach/include
//
// #include <stdlib.h>
// #include <libroach.h>
import "C"

func main() {
	fmt.Println("sup losers!")
}
