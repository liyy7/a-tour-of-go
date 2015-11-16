package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	messages := make(chan int)
	done := make(chan interface{})
	var wg sync.WaitGroup

	// Receive numbers from message channel
	go func() {

		fmt.Println("RECEIVE IN")

		// Receive and print number until a DONE signal got
		func() {
			for {
				select {
				case msg := <-messages:
					fmt.Println(msg, "RECEIVED")
				case <-done:
					fmt.Println("DONE SIGNAL RECEIVED")
					return
				}
			}
		}()

		/*
		 * The `for` statement won't finish
		 *
		 */
		//		for msg := range messages {
		//			fmt.Println(msg, "RECEIVED")
		//		}

		fmt.Println("RECEIVE OUT")
	}()

	// Get the loop number from STDIN
	fmt.Println("Input a number...")
	var N int
	if _, err := fmt.Scanf("%d", &N); err != nil {
		fmt.Errorf("INPUT ERROR:", err)
	}

	// Send `N` numbers to message channel
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(N-n))
			fmt.Println(n, "SENDING")
			messages <- n
		}(i + 1)
	}

	// Wait for all groups
	wg.Wait()

	// Send DONE signal
	fmt.Println("DONE SIGNAL SENDING")
	done <- struct{}{}
}
