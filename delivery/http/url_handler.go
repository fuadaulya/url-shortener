package http

import (
	"context"
	"encoding/json"
	"net/http"

	"url-shortener-db-migrate/pkg/entity"
	errors "url-shortener-db-migrate/pkg/entity"

	"github.com/julienschmidt/httprouter"
)

// CreateShortURL handles HTTP POST request for creating a new short URL.
func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	var url entity.URL

	// Parse JSON body request
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Call the usecase to handle business logic
	shortenedURL, err := h.usecase.CreateShortURL(ctx, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the shortened URL as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortenedURL)
}

// GetURLByShort handles GET requests for retrieving URLs by their short version using query parameter.
func (h *URLHandler) GetURLByShort(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	// Retrieve the short URL from query parameters
	shortURL := r.URL.Query().Get("shortURL")

	// Validate the short URL parameter
	if shortURL == "" {
		http.Error(w, "shortURL query parameter is required", http.StatusBadRequest)
		return
	}

	// Get the original URL from usecase
	url, err := h.usecase.GetURLTargetByShort(ctx, shortURL)
	if err != nil {
		if err == errors.ErrNotFound {
			http.Error(w, "URL not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, url.URLTarget, http.StatusFound) // or http.StatusMovedPermanently for 301
}

// GetAllURLs handles GET requests to retrieve all URLs.
func (h *URLHandler) GetAllURLs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	urls, err := h.usecase.GetAllURLs(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(urls)
}

// UpdateURLShort handles PUT requests to update a short URL.
func (h *URLHandler) UpdateURLShort(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	var url entity.URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.usecase.UpdateURLShort(ctx, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateURLShort handles PUT requests to update a target URL.
func (h *URLHandler) UpdateURLTarget(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	var url entity.URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.usecase.UpdateURLTarget(ctx, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
