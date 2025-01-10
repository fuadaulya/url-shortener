package usecase

import (
	"context"
	"url-shortener-db-migrate/pkg/entity"
)

// URLRepository adalah interface yang mendefinisikan metode yang diperlukan untuk berinteraksi dengan data URL.
type URLRepository interface {
	InsertURL(ctx context.Context, url entity.URL) (id int, err error)
	UpdateURLShort(ctx context.Context, url entity.URL) error
	UpdateURLTarget(ctx context.Context, url entity.URL) error
	GetURLTargetByURLShort(ctx context.Context, urlInput entity.URL) (entity.URL, error)
	GetAllURL(ctx context.Context) ([]entity.URL, error)
}
