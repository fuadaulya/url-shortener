package main

/*
import (
	"log"
	"net/http"
	delivery "url-shortener-db-migrate/delivery/http"
	"url-shortener-db-migrate/pkg/database"
	"url-shortener-db-migrate/repository"
	"url-shortener-db-migrate/usecase"
)

func main() {
	// ctx := context.Background()

	// Initialize database
	repository.InitDatabase()
	defer database.CloseDB()
	//_________________

	// Inisialisasi repository
	repoPostgre := repository.NewURLRepository()

		// Get user by ID
		urlTarget, err := repoPostgre.GetURLTargetByURLShort(ctx, entity.URL{
			URLShort: "short.url/fuad-twitter",
		})
		if err != nil {
			log.Fatalf("Error getting url target by url short: %v", err)
		}
		log.Printf("URLTarget: %#v\n", urlTarget)

		// Get all users
		urls, err := repoPostgre.GetAllURL(ctx)
		if err != nil {
			log.Fatalf("Error getting all urls: %v", err)
		}
		log.Printf("URLs: %#v\n", urls)
	//_________________

	// Inisialisasi usecase
	uc := usecase.NewUsecase(repoPostgre)
		// Contoh penggunaan metode dari usecase

		// Memanggil metode GetURLTargetByShort dari usecase
		urlShort := "short.url/example" // Ganti dengan URL short yang relevan
		url, err := uc.URLUsecase.GetURLTargetByShort(ctx, urlShort)
		if err != nil {
			log.Fatalf("Error getting URL target by short: %v", err)
		}

		// Menampilkan hasil
		log.Printf("URL target for '%s': %s\n", urlShort, url.URLTarget)
	//_________________

	// Initialize router
	router := delivery.NewRouter(uc.URLUsecase)

	// Start HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	//_________________

	// fmt.Println("Hello, world!")
	log.Println("Application is running...")
}
*/
