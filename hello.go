package main

import (
	"fmt"
	"unicode/utf8"
)

type String string

func (s String) Take(idx int) String {
	runes := []rune(string(s))
	if idx < len(runes) {
		return String(runes[idx])
	} else {
		return String("")
	}
}

func main() {
	str := String("hello, 世界")
	for i := 0; i <= utf8.RuneCountInString(string(str)); i++ {
		fmt.Print(str.Take(i))
	}
	fmt.Println()
}
