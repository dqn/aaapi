package main

import (
	"log"
	"net/http"

	"github.com/dghubble/oauth1"
)

func newClient(ck, cs, at, as string) *http.Client {
	config := oauth1.NewConfig(ck, cs)
	token := oauth1.NewToken(at, as)
	client := config.Client(oauth1.NoContext, token)

	return client
}

func run() error {
	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
