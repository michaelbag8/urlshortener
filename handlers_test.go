package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateURLHandler(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{
			name:       "valid request",
			body:       `{"original_url":"https://google.com"}`,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "invalid json",
			body:       `{"original_url":`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "empty url",
			body:       `{"original_url":""}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPost,
				"/urls",
				strings.NewReader(tt.body),
			)

			w := httptest.NewRecorder()

			createURLHandler(w, req)

			res := w.Result()

			if res.StatusCode != tt.wantStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.wantStatus,
					res.StatusCode,
				)
			}
		})
	}
}

func TestCreateURLHandler_Success(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/urls",
		strings.NewReader(`{"original_url":"https://google.com"}`),
	)

	w := httptest.NewRecorder()

	createURLHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf(
			"expected %d got %d",
			http.StatusCreated,
			res.StatusCode,
		)
	}

	var response URL

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.OriginalURL != "https://google.com" {
		t.Fatalf(
			"expected original url %q got %q",
			"https://google.com",
			response.OriginalURL,
		)
	}

	if response.ShortCode == "" {
		t.Fatal("expected shortcode to be generated")
	}

	if response.Clicks != 0 {
		t.Fatalf(
			"expected clicks 0 got %d",
			response.Clicks,
		)
	}
}

func TestRedirectHandler_NotFound(t *testing.T) {
	urlStore = make(map[string]URL)

	req := httptest.NewRequest(
		http.MethodGet,
		"/missing",
		nil,
	)

	req.SetPathValue("code", "missing")

	w := httptest.NewRecorder()

	redirectHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
			res.StatusCode,
		)
	}
}