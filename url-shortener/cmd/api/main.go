package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	app "github.com/mohamed-gudle/backend-projects/url-shortener/internal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}

	application := app.NewApplication(db)

	server := &http.Server{
		Addr:     ":4000",
		Handler:  application.Routes(),
	}

	fmt.Println("Starting server on :4000")
	server.ListenAndServe()
	fmt.Println("Server stopped")
}
