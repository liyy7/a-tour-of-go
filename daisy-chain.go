package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	n := 100000
	leftmost := make(chan int)
	var left, right chan int
	left = leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
