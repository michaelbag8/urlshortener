package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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

	urlStore[shortCode] = result

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	url, exist := urlStore[code]
	if !exist {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	url.Clicks++
	urlStore[code] = url

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		log.Printf("error encoding response: %v", err)
	}
}
