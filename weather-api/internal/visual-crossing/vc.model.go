package vc

type Weather struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ResolvedAddress string `json:"resolvedAddress"`
	Address string `json:"address"`
	Description string `json:"description"`
	days []Day `json:"days"`
}

type Day struct {
	Date string `json:"datetime"`
	Temperture float64 `json:"temp2m"`
	feelsLike float64 `json:"feelslike"`

}


