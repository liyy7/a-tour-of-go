package main

import (
	"fmt"
)

type Man struct {
	name string
}

func (m Man) setNameWithValue(name string) {
	fmt.Printf("Calling %q.setNameWithValue(%q)\n", m, name)
	m.name = name
}

func (m *Man) setNameWithPointer(name string) {
	fmt.Printf("Calling %q.setNameWithPointer(%q)\n", m, name)
	m.name = name
}

func main() {
	man := Man{name: "John"}

	fmt.Printf("Before setNameWithValue: %q\n", man)
	man.setNameWithValue("Green")
	fmt.Printf("After setNameWithValue: %q\n", man)

	fmt.Println()

	fmt.Printf("Before setNameWithPointer: %q\n", man)
	(&man).setNameWithPointer("Green")
	fmt.Printf("After setNameWithPointer: %q\n", man)
}
