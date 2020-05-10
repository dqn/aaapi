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
	r, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// println(string(b))

	var resp ErrorResponse
	json.Unmarshal(b, &resp)
	if errs := resp.Errors; len(errs) != 0 {
		err = fmt.Errorf("code %d: %s", errs[0].Code, errs[0].Message)
		return nil, err
	}

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

	var resp PostWebhooksResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
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

	var resp GetWebhooksResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
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

	var resp GetWebhooksWithEnvNameResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (a *AAAPI) PutWebhooks(webhookID string) (*PutWebhooksResponse, error) {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/webhooks/%s.json", a.envName, webhookID)
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}

	_, err = a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var resp PutWebhooksResponse

	return &resp, nil
}

func (a *AAAPI) PostSubscriptions() (*PostSubscriptionsResponse, error) {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/subscriptions.json", a.envName)
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}

	_, err = a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var resp PostSubscriptionsResponse

	return &resp, nil
}

func (a *AAAPI) GetSubscriptionsCount() (*GetSubscriptionsCountResponse, error) {
	u := *a.url
	u.Path += "/all/subscriptions/count.json"
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	b, err := a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var resp GetSubscriptionsCountResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (a *AAAPI) GetSubscriptions() (*GetSubscriptionsResponse, error) {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/subscriptions.json", a.envName)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	_, err = a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var resp GetSubscriptionsResponse

	return &resp, nil
}

func (a *AAAPI) GetSubscriptionsList() (*GetSubscriptionsListResponse, error) {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/subscriptions/list.json", a.envName)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	b, err := a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var resp GetSubscriptionsListResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (a *AAAPI) DeleteWebhooks(webhookID string) (*DeleteWebhooksResponse, error) {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/webhooks/%s.json", a.envName, webhookID)
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}

	_, err = a.prosessRequest(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteWebhooksResponse

	return &resp, nil
}
