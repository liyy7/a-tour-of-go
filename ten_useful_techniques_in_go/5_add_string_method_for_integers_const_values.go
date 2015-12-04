package main

import (
	"fmt"
)

type State int

const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

//go:generate stringer -type=State

func main() {
	state := Running
	fmt.Println("state ", state)
}
