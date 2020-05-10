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
	r, err := a.PutWebhooks(os.Getenv("WEBHOOK_ID"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}

func TestPostSubscriptions(t *testing.T) {
	a := newAAAPI()
	r, err := a.PostSubscriptions()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}

func TestGetSubscriptionsCount(t *testing.T) {
	a := newAAAPI()
	r, err := a.GetSubscriptionsCount()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}

func TestGetSubscriptions(t *testing.T) {
	a := newAAAPI()
	r, err := a.GetSubscriptions()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}

func TestGetSubscriptionsList(t *testing.T) {
	a := newAAAPI()
	r, err := a.GetSubscriptionsList()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}

func TestDeleteWebhooks(t *testing.T) {
	a := newAAAPI()
	r, err := a.DeleteWebhooks(os.Getenv("WEBHOOK_ID"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}

func TestDeleteSubscriptionsUser(t *testing.T) {
	a := newAAAPI()
	r, err := a.DeleteSubscriptionsUser(os.Getenv("USER_ID"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}
