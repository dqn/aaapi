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
	fmt.Printf("%+v\n", r)
}

func TestGetWebhooks(t *testing.T) {
	a := newAAAPI()
	r, err := a.GetWebhooks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}

func TestGetWebhooksWithEnvName(t *testing.T) {
	a := newAAAPI()
	r, err := a.GetWebhooksWithEnvName(os.Getenv("ENV"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}

func TestPutWebhooks(t *testing.T) {
	a := newAAAPI()
	err := a.PutWebhooks(os.Getenv("WEBHOOK_ID"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPostSubscriptions(t *testing.T) {
	a := newAAAPI()
	err := a.PostSubscriptions()
	if err != nil {
		t.Fatal(err)
	}
}
