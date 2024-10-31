package vc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type vcController struct {
	service IVCService
}

func NewVCController(service IVCService) *vcController {
	return &vcController{service: service}
}

func(v *vcController) GetWeather( Response http.ResponseWriter,Request *http.Request) {
	location := Request.URL.Query().Get("location")
	weather, err := v.service.GetWeather(location)

	fmt.Printf("Weather: %+v\n", weather)

	if err != nil {
		Response.Write([]byte(err.Error()))
		return
	}

	data,err:=json.Marshal(weather)

	if err != nil {
		Response.Write([]byte(err.Error()))
		return
	}


	Response.Write(data)
}