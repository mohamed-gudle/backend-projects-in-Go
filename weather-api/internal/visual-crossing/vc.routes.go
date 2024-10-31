package vc

import "net/http"

func VCRoutes() *http.ServeMux {
	vcRepository := NewRedisVCRepository()
	vcService := NewVCService(vcRepository)
	vcController := NewVCController(vcService)


	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/weather", vcController.GetWeather)
	return mux
}
