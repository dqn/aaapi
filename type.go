package aaapi

import "time"

type PostWebhooksResponse struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
}

type GetWebhooksResponse struct {
	Environments []Environments `json:"environments"`
}

type Webhooks struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
}

type Environments struct {
	EnvironmentName string     `json:"environment_name"`
	Webhooks        []Webhooks `json:"webhooks"`
}

type GetWebhooksWithEnvNameResponse []Webhooks
