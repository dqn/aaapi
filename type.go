package aaapi

import "time"

type PostWebhooksResponse struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
}
