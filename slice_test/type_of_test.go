package test_slices

import (
	"fmt"
	"testing"
)

func TestTypeOf(t *testing.T) {
	cases := []struct {
		in   interface{}
		want string
	}{
		{[]int{1, 2, 3, 5, 8, 11}, "[]int"},
		{[6]int{1, 2, 3, 5, 8, 11}, "[6]int"},   // explicit
		{[...]int{1, 2, 3, 5, 8, 11}, "[6]int"}, //implicit
	}
	for _, c := range cases {
		got := fmt.Sprintf("%T", c.in)
		if got != c.want {
			t.Errorf("TypeOf(%#v) == %q, want %q", c.in, got, c.want)
		} else {
			t.Logf("TypeOf(%#v) == %q", c.in, got)
		}
	}
}
