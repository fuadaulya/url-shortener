package http

import (
	"context"
	"url-shortener-db-migrate/pkg/entity"
)

// Definisi interface untuk usecase
type URLUsecase interface {
	// ShortenURL menambahkan URL short baru
	CreateShortURL(ctx context.Context, url entity.URL) (entity.URL, error)

	// GetURLTargetByShort mendapatkan target URL berdasarkan URL short
	GetURLTargetByShort(ctx context.Context, urlShort string) (entity.URL, error)

	// GetAllURLs mendapatkan semua URL
	GetAllURLs(ctx context.Context) ([]entity.URL, error)

	// UpdateURLShort mengubah short URL yang ada
	UpdateURLShort(ctx context.Context, url entity.URL) error

	// UpdateURLTarget mengubah target URL yang ada
	UpdateURLTarget(ctx context.Context, url entity.URL) error
}
