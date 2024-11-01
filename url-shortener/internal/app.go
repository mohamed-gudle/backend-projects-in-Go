package app

import (
	"net/http"

	shorturl "github.com/mohamed-gudle/backend-projects/url-shortener/internal/short-url"
	"gorm.io/gorm"
)

type IApplication interface {
	*gorm.DB
}

type application struct {
	DB *gorm.DB
}

func NewApplication(db *gorm.DB) *application{
	return &application{
		DB:db,
	}
}

func (app *application) Routes () *http.ServeMux {
	mux:= http.NewServeMux()
	mux.Handle("/",shorturl.Routes(app.DB))
	return mux
}