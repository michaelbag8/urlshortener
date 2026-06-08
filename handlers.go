package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}


func createURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req struct {
		OriginalURL string `json:"original_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.OriginalURL == "" {
		http.Error(w, "Empty field", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()

	result := URL{
		OriginalURL: req.OriginalURL,
		Clicks:      0,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().AddDate(0, 0, 30),
		ShortCode:   shortCode,
	}

	mu.Lock()
	urlStore[shortCode] = result
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	mu.RLock()
	url, exist := urlStore[code]
	mu.RUnlock()
	if !exist {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	mu.Lock()
	url.Clicks++
	urlStore[code] = url
	mu.Unlock()

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		log.Printf("error encoding response: %v", err)
	}
}
