package main

import "fmt"

// #cgo CPPFLAGS: -Ic-deps/libroach/include
// #cgo LDFLAGS: -L/Users/irfansharif/Software/native/x86_64-apple-darwin19.6.0/libroach -lroach
//
// #include <stdlib.h>
// #include <libroach.h>
import "C"

func main() {
	fmt.Println("sup losers!")
}
