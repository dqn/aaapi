package aaapi

import (
	"fmt"
	"os"
	"testing"
)

func newAAAPI() *AAAPI {
	return NewPremium(
		os.Getenv("CK"),  // consumer key
		os.Getenv("CS"),  // consumer secret
		os.Getenv("AT"),  // access token
		os.Getenv("AS"),  // access token secret
		os.Getenv("ENV"), // env name
	)
}

func TestPostWebhooks(t *testing.T) {
	a := newAAAPI()
	r, err := a.PostWebhooks(os.Getenv("CRC"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", r)
}

func TestGetWebhooks(t *testing.T) {
	a := newAAAPI()
	r, err := a.GetWebhooks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", r)
}
