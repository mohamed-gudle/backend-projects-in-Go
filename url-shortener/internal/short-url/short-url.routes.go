package shorturl

import (
	"net/http"
	"gorm.io/gorm"
)



func Routes(db *gorm.DB) *http.ServeMux {
	repository := NewShortURLRepository(db)
	service := NewDefaultShortURLService(repository)
	controller := NewShortURLController(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /short-url", controller.CreateShortURL)
	return mux
}