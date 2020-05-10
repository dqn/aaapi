package aaapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

const baseURL = "https://api.twitter.com/1.1/account_activity"

type AAAPI struct {
	client  *http.Client
	url     *url.URL
	envName string
}

func NewPremium(ck, cs, at, as, envName string) *AAAPI {
	config := oauth1.NewConfig(ck, cs)
	token := oauth1.NewToken(at, as)
	client := config.Client(oauth1.NoContext, token)
	u, _ := url.Parse(baseURL)

	return &AAAPI{
		client:  client,
		envName: envName,
		url:     u,
	}
}

func (a *AAAPI) prosessRequest(req *http.Request) ([]byte, error) {
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// println(string(b))

	return b, nil
}

func (a *AAAPI) PostWebhooks(rawURL string) (*PostWebhooksResponse, error) {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/webhooks.json", a.envName)
	q := url.Values{
		"url": {rawURL},
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}

	b, err := a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var r PostWebhooksResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (a *AAAPI) GetWebhooks() (*GetWebhooksResponse, error) {
	u := *a.url
	u.Path += "/all/webhooks.json"

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	b, err := a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var r GetWebhooksResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (a *AAAPI) GetWebhooksWithEnvName(envName string) (*GetWebhooksWithEnvNameResponse, error) {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/webhooks.json", envName)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	b, err := a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var r GetWebhooksWithEnvNameResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (a *AAAPI) PutWebhooks(webhookID string) error {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/webhooks/%s.json", a.envName, webhookID)
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
