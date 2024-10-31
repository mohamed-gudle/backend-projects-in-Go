package main

import "net/http"
import "github.com/mohamed-gudle/backend-projects/weather-api/internal/visual-crossing"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	mux.Handle("/v1/", vc.VCRoutes())
	return mux
}