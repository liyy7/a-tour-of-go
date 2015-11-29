package main

import (
	"a-tour-of-go/local_package_test/string"
	"fmt"
)

func main() {
	var s string.String = "hello,世界"
	fmt.Printf("len(%q) = %d\n", s, len(s))
	fmt.Printf("%q.Length() = %d\n", s, s.Length())
}
