package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan interface{})
	go func() {
		select {
		case m := <-ch:
			fmt.Println(m)
		case <-time.After(2 * time.Second):
			fmt.Println("time out")
		}
	}()

	time.Sleep(3 * time.Second)
}
