package vc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

type IVCService interface {
	GetWeather(location string) (*Weather, error)
}

type VCService struct {
	repository IVCRepository
	apiKey     string
}

func NewVCService(respository IVCRepository) *VCService {
	return &VCService{repository: respository,
		apiKey: os.Getenv("WEATHER_API_KEY")}
}

var ctx = context.Background()

func (v *VCService) GetWeather(location string) (*Weather, error) {

	weatherString, err := v.repository.GetWeather(location)

	if err == redis.Nil {

		response, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + location + "?key=" + v.apiKey)

		if err != nil {
			return nil, err
		}

		fmt.Printf("Response: %+v\n", response)
		data, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		var weather Weather

		json.Unmarshal(data, &weather)

		v.repository.SetWeather(location, string(data))

		return &weather, nil

	}

	if err != nil {
		return nil, err
	}

	var weather Weather
	json.Unmarshal([]byte(weatherString), &weather)

	return &weather, nil
}
