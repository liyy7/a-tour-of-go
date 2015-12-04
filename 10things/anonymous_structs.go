package main

import (
	"fmt"
	"sync"
)

// Grouped globals
var config struct {
	APIKey    string
	APISecret string
}

func main() {
	config.APIKey = "BADC0C0A"
	config.APISecret = "********"

	fmt.Printf("Grouped globbals, data: %q\n", config)

	// Template data
	type User struct {
		Name      string
		BirthYear int
	}
	title := "Programming languages"
	users := []*User{
		{Name: "Go", BirthYear: 2009},
		{Name: "Ruby", BirthYear: 1995},
	}
	data := struct {
		Title string
		Users []*User
	}{
		Title: title,
		Users: users,
	}

	fmt.Printf("Template data, data: %q\n", data)

	// Test tables
	var indexRuneTests = []struct {
		s    string
		rune rune
		out  int
	}{
		{s: "a A x", rune: 'A', out: 2},
		{s: "some_text=some_value", rune: '=', out: 9},
		{s: "☺a", rune: 'a', out: 3},
		{s: "a☻☺b", rune: '☺', out: 4},
	}

	fmt.Printf("Test tables, indexRuneTest: %q\n", indexRuneTests)

	// Embedded lock
	var hits struct {
		sync.Mutex
		n int
	}
	hits.Lock()
	hits.n++
	hits.Unlock()

	fmt.Printf("Embedded lock, hits: %q\n", hits)
}
