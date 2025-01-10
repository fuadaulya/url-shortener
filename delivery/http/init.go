/*
	- Inisialisasi HTTP handler dengan usecase
	- Inisialisasi HTTP handler dengan dependency injection (usecase)
*/

package http

import (
	"github.com/julienschmidt/httprouter"
)

type URLHandler struct {
	usecase URLUsecase
}

func NewURLHandler(uc URLUsecase) *URLHandler {
	return &URLHandler{
		usecase: uc,
	}
}

// NewRouter initializes the HTTP router and routes the requests.
func NewRouter(uc URLUsecase) *httprouter.Router {
	r := httprouter.New()

	handler := NewURLHandler(uc)

	r.POST("/shorten", handler.CreateShortURL)
	r.GET("/short", handler.GetURLByShort)
	r.GET("/urls", handler.GetAllURLs)
	r.PUT("/update-short", handler.UpdateURLShort)
	r.PUT("/update-target", handler.UpdateURLTarget)

	return r
}
