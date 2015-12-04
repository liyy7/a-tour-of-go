package main

import (
	"fmt"
)

type State int

const (
	Unknown State = iota
	Running
	Stopped
	Rebooting
	Terminated
)

//go:generate stringer -type=State

func main() {
	state := Running
	fmt.Println("state ", state)

	type T struct {
		Name  string
		Port  int
		State State
	}

	t := T{Name: "example", Port: 6666}

	fmt.Printf("t %+v\n", t)
}
