package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	const responseSample = `
		{"data": {"children": [
		  {"data": {
			"title": "The Go homepage",
			"url": "http://golang.org/"
		  }},
		  {"data": {
			"title": "The Ruby homepage",
			"url": "https://www.ruby-lang.org/"
		  }}
		]}}
		`

	type Item struct {
		Title string
		URL   string
	}
	type Response struct {
		Data struct {
			Children []struct {
				Data Item
			}
		}
	}

	decoder := json.NewDecoder(strings.NewReader(responseSample))
	var response Response
	if err := decoder.Decode(&response); err != nil {
		fmt.Printf("Nested structs, err: %q", err)
	} else {
		fmt.Printf("Nested structs, response: %q", response)
	}
}
