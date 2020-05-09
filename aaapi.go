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

func (a *AAAPI) PostWebhooks(rawURL string) (*PostWebhooksResponse, error) {
	u := *a.url
	u.Path += fmt.Sprintf("/all/%s/webhooks.json", a.envName)
	q := url.Values{
		"url": {url.QueryEscape(rawURL)},
	}
	u.RawQuery = q.Encode()

	resp, err := a.client.Post(u.String(), "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r PostWebhooksResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
