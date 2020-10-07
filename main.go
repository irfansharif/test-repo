package main

import "fmt"

// #cgo CXXFLAGS: -std=c++14
// #cgo CPPFLAGS: -Ic-deps/googletest/googletest/include/gtest
//
// #include <gtest.h>
import "C"

func main() {
	fmt.Println("vim-go")
}
