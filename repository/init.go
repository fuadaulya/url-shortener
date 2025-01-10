package repository

import (
	"database/sql"
	"url-shortener-db-migrate/config"
	db "url-shortener-db-migrate/pkg/database"
)

type Postgres struct {
	DB *sql.DB
}

func InitDatabase() {
	cfg := config.GetConfig()
	db.InitDB(cfg)
}

// NewURLRepository menginisialisasi URLRepository dengan koneksi database.
func NewURLRepository() *Postgres {
	return &Postgres{
		DB: db.DB,
	}
}
