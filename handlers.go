package main

import (
	"net/http"
	"time"
	"encoding/json"
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

	if req.OriginalURL == ""{
		http.Error(w, "Empty field", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()

	result := URL{
		OriginalURL: req.OriginalURL,
		Clicks: 0,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().AddDate(0,0,30),
		ShortCode: shortCode,
	}


	urlStore[shortCode] = result

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)


}