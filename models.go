package main

import (
	"time"
)

type URL struct {
	OriginalURL string `json:"original_url"`
	Clicks      int		`json:"clicks"`
	CreatedAt   time.Time	`json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	ShortCode   string	`json:"short_code"`
}

