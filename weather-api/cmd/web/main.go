package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type application struct{
	logger log.Logger
	redisClient *redis.Client
}

var ctx = context.Background()

func main() {
	godotenv.Load()


	application := &application{
		logger: *log.New(os.Stdout, "", log.Ldate|log.Ltime),
		redisClient: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		}),
		
	}

	server := &http.Server{
		Addr:    ":4000",
		Handler: application.routes(),
		ErrorLog: &application.logger,

	}

	application.logger.Println("Starting server on :4000")
	err:=server.ListenAndServe()

	if err != nil {
		application.logger.Println("Error starting server: ", err)
	}
	application.logger.Println("Server stopped")

}