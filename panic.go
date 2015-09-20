package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type MyError struct {
	string
}

func (err MyError) Error() string {
	return err.string
}

type UnknownError struct {
	string
}

func foo() (err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			fmt.Printf("[foo] %v found: %v\n", reflect.TypeOf(recoverErr), recoverErr)
			if myError, found := recoverErr.(MyError); found {
				err = myError
			}
		}
	}()

	if i := rand.Intn(2); i != 0 {
		panic(MyError{string: fmt.Sprintf("panic here (%v)", i)})
	} else {
		panic(UnknownError{fmt.Sprintf("unknown error (%v)", i)})
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := foo(); err != nil {
		fmt.Printf("[main] %v found: %v\n", reflect.TypeOf(err), err)
	} else {
		fmt.Println("[main] No error found")
	}
}
