package main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.Kitchen))
	// https://golang.org/pkg/time/#Time.Format
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(time.Now().Format("01/02 15:04"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("06/1/2 15:4:5.0 -0700"))
	fmt.Println(time.Now().Format("06/1/2 15:4:5.00 -0700"))
	fmt.Println(time.Now().Format("06/1/2 15:4:5.000 -0700"))
	fmt.Println(time.Now().Unix())
}