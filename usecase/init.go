package usecase

// import (
//
//	"url-shortener-db-migrate/repository"
//
// )

type URLUsecase struct {
	repo URLRepository
}

// NewURLUsecase creates a new instance of URLUsecase with the given repository.
func NewURLUsecase(repo URLRepository) *URLUsecase {
	return &URLUsecase{repo: repo}
}
