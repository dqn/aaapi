package aaapi

import (
	"time"
)

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type PostWebhooksResponse struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
}

type GetWebhooksResponse struct {
	Environments []Environment `json:"environments"`
}

type Webhook struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
}

type Environment struct {
	EnvironmentName string    `json:"environment_name"`
	Webhooks        []Webhook `json:"webhooks"`
}

type GetWebhooksWithEnvNameResponse []Webhook

type PutWebhooksResponse struct{}

type PostSubscriptionsResponse struct{}

type GetSubscriptionsCountResponse struct {
	AccountName        string `json:"account_name"`
	SubscriptionsCount string `json:"subscriptions_count"`
	ProvisionedCount   string `json:"provisioned_count"`
}

type GetSubscriptionsResponse struct{}
