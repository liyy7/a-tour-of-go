package main

import (
	"fmt"
)

type Man struct {
	name string
}

func (m *Man) creep() {
	fmt.Printf("%s is creeping\n", m.name)
}
func (m *Man) walk() {
	fmt.Printf("%s is walking\n", m.name)
}
func (m *Man) run() {
	fmt.Printf("%s is running\n", m.name)
}

func main() {
	type Fn func(*Man)
	m := Man{name: "John"}
	for _, fn := range []Fn{(*Man).creep, (*Man).walk, (*Man).run} {
		fn(&m)
	}
}
