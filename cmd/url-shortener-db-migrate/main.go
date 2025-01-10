package main

import (
	"log"
	"net/http"
	delivery "url-shortener-db-migrate/delivery/http"
	"url-shortener-db-migrate/pkg/database"
	"url-shortener-db-migrate/repository"
	"url-shortener-db-migrate/usecase"

	"github.com/joho/godotenv"
)

func main() {
	// read .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed to load .env")
	}
	//________________

	// ctx := context.Background()

	// Initialize database
	repository.InitDatabase()
	defer database.CloseDB()
	//_________________

	// Inisialisasi repository
	repoPostgre := repository.NewURLRepository()

	// Inisialisasi usecase
	uc := usecase.NewURLUsecase(repoPostgre)

	// Initialize router
	router := delivery.NewRouter(uc)

	// Start HTTP server
	log.Println("Starting server on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	//_________________

	// fmt.Println("Hello, world!")
	log.Println("Application is running...")
}
