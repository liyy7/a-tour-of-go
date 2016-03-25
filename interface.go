package main

import (
	"fmt"
	"reflect"
)

type Human interface {
	Sleep()
	Eat()
}

type Asian struct {
	name string
}

func (a Asian) String() string {
	return a.name
}

func (a *Asian) Sleep() {
	fmt.Printf("%q is sleeping\n", a)
}

func (a *Asian) Eat() {
	fmt.Printf("%q is eating\n", a)
}

func main() {
	var a Human = &Asian{
		name: "Bejito",
	}
	a.Sleep()
	a.Eat()
	fmt.Println(a, reflect.TypeOf(a.(*Asian)), reflect.TypeOf(a.(Human)))

}
