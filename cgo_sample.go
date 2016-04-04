package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"time"
)

func Seed(i int64) {
	C.srandom(C.uint(i))
}

func Random() int {
	return int(C.random())
}

func main() {
	Seed(time.Now().UnixNano())
	fmt.Println(Random())
}
