package usecase

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"strings"
	"url-shortener-db-migrate/pkg/entity"

	"github.com/google/uuid"
)

// CreateShortURL adalah logika bisnis untuk menambahkan URL baru
func (u *URLUsecase) CreateShortURL(ctx context.Context, url entity.URL) (entity.URL, error) {
	// Validasi URL short dan URL target
	if url.URLTarget == "" {
		return entity.URL{}, errors.New("URL short and URL target cannot be empty")
	}

	// Cek apakah URL short sudah ada
	existingURL, err := u.repo.GetURLTargetByURLShort(ctx, entity.URL{
		URLShort: url.URLShort,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return entity.URL{}, err
	}

	if existingURL.URLShort != "" {
		return entity.URL{}, errors.New("URL short already exists")
	}

	// Generate short URL
	url.URLShort = GenerateShortURL(existingURL.URLTarget)

	// Insert URL baru ke database
	id, err := u.repo.InsertURL(ctx, url)
	if err != nil {
		return entity.URL{}, err
	}

	url.ID = id

	return url, nil
}

// GetURLTargetByShort adalah logika bisnis untuk mendapatkan URL target berdasarkan URL short
func (u *URLUsecase) GetURLTargetByShort(ctx context.Context, urlShort string) (entity.URL, error) {
	if urlShort == "" {
		return entity.URL{}, errors.New("URL short cannot be empty")
	}

	// Dapatkan URL target dari repository
	url, err := u.repo.GetURLTargetByURLShort(ctx, entity.URL{
		URLShort: urlShort,
	})
	if err != nil {
		return entity.URL{}, err
	}

	// Cek apakah URL ditemukan
	if url.URLShort == "" {
		return entity.URL{}, errors.New("URL not found")
	}

	return url, nil
}

// GetAllURLs adalah logika bisnis untuk mendapatkan semua URL yang ada
func (u *URLUsecase) GetAllURLs(ctx context.Context) ([]entity.URL, error) {
	// Dapatkan semua URL dari repository
	urls, err := u.repo.GetAllURL(ctx)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

// UpdateURLShort adalah logika bisnis untuk memperbarui short URL
func (u *URLUsecase) UpdateURLShort(ctx context.Context, url entity.URL) error {
	// Validasi URL short
	if url.URLShort == "" {
		return errors.New("URL short and URL target cannot be empty")
	}

	// Update URL di database
	return u.repo.UpdateURLShort(ctx, url)
}

// UpdateURL adalah logika bisnis untuk memperbarui target URL
func (u *URLUsecase) UpdateURLTarget(ctx context.Context, url entity.URL) error {
	// Validasi URL target
	if url.URLTarget == "" {
		return errors.New("URL short and URL target cannot be empty")
	}

	// Update URL di database
	return u.repo.UpdateURLTarget(ctx, url)
}

// GenerateShortURL generates a short URL using a hash of the original URL.
func GenerateShortURL(originalURL string) string {
	// Create a SHA-256 hash of the original URL + UUID
	hash := sha256.New()
	uniqueID := uuid.New().String()
	hash.Write([]byte(originalURL + uniqueID))

	// Get the first 6 characters of the hash as the short URL
	hashValue := hex.EncodeToString(hash.Sum(nil))

	// Convert the first 6 characters to a more URL-friendly format (lowercase)
	shortURLPart := strings.ToLower(hashValue[:6])

	baseShortURL := "short.url/"

	return baseShortURL + shortURLPart
}
