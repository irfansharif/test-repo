package main

import "fmt"

// #cgo CPPFLAGS: -Ic-deps/libroach/include
// #cgo CPPFLAGS: -Ic-deps/rocksdb/include
//
// #include <libroach.h>
// #include <rocksdb/db.h>
import "C"

func main() {
	fmt.Println("sup losers!")
}
