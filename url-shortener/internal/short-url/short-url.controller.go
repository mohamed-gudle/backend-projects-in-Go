package shorturl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IShortURLController interface {
	CreateShortURL(r *http.Request, w http.ResponseWriter)
	GetShortURL(r *http.Request, w http.ResponseWriter)
	UpdateShortURL(r *http.Request, w http.ResponseWriter)
	DeleteShortURL(r *http.Request, w http.ResponseWriter)
}

type shortURLController struct {
	service IShortURLService
}

func NewShortURLController(service IShortURLService) *shortURLController {
	return &shortURLController{service: service}
}

func (c *shortURLController) CreateShortURL(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))

	}
	var args struct {
		ShortCode string `json:"short_code"`
		URL       string `json:"url"`
	}

	fmt.Println(body)

	fmt.Println("Read Body")

	err = json.Unmarshal(body, &args)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Printf(`%+v`, args)

	err = c.service.CreateShortURL(args.ShortCode, args.URL)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

func (c *shortURLController) GetShortURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Query().Get("short_code")
	shortURL, err := c.service.GetShortURL(shortCode)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	data, err := json.Marshal(shortURL)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}

func (c *shortURLController) UpdateShortURL(w http.ResponseWriter, r *http.Request) {
	var updatedURL struct {
		Url string `json:"url"`
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(body, &updatedURL)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	shortCode := r.URL.Query().Get("short_code")

	err = c.service.UpdateShortURL(shortCode, updatedURL.Url)

	if err != nil {
		w.Write([]byte(err.Error()))
		return

	}

}

func (c *shortURLController) DeleteShortURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Query().Get("short_code")

	err := c.service.DeleteShortURL(shortCode)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

