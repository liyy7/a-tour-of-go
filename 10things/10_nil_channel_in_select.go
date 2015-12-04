package main

import (
	"fmt"
	"time"
)

type Work struct {
	number int
}

func (w Work) String() string {
	return fmt.Sprintf("work_%d", w.number)
}
func (w Work) Refuse() {}
func (w Work) Do()     {}

func worker(i int, ch chan Work, quit chan struct{}) {
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case w := <-ch:
			if quit == nil {
				w.Refuse()
				fmt.Println("worker", i, "refused", w)
				return
			}
			w.Do()
			fmt.Println("worker", i, "processed", w)
		case <-quit:
			fmt.Println("worker", i, "quitting")
			quit = nil
		}
	}
}

func main() {
	ch, quit := make(chan Work), make(chan struct{})
	go func() {
		for i := 0; ; i++ {
			ch <- Work{number: i}
		}
	}()
	for i := 0; i < 4; i++ {
		go worker(i, ch, quit)
	}
	time.Sleep(5 * time.Second)
	close(quit)
	time.Sleep(2 * time.Second)
}
