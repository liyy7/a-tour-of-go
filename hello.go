package main

import (
	"fmt"
)

type String string

func (s String) Take(idx int) String {
	runes := []rune(string(s))
	if idx < len(runes) {
		return String(runes[idx])
	} else {
		return ""
	}
}

func main() {
	str := String("hello, 世界")
	for i := range []rune(string(str)) {
		fmt.Print(str.Take(i))
	}
	fmt.Println()
}
