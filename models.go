package main

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


type URL struct {
	OriginalURL string
	Clicks      int
	CreatedAt   time.Time
	ExpiresAt   time.Time
	ShortCode   string
}

func generateShortCode() string {
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
