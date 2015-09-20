package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	conf := &oauth2.Config{
		ClientID:     "client_id",
		ClientSecret: "client_secret",
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://example.com/auth",
			TokenURL: "http://example.com/token",
		},
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit this URL for auth dialog: %v\n", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}
	client := conf.Client(oauth2.NoContext, tok)
	client.Get("http://...")
}
