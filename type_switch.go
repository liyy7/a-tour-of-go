package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var values []interface{} = []interface{}{0, "anything", true}
	var anything interface{} = values[rand.Intn(len(values))]
	switch v := anything.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Printf("default(%v): %v\n", reflect.TypeOf(v), v)
	}
}
