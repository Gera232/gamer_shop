package main

import (
	"back-end/api"
	"back-end/storage"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	api.SetupRoutes()

	storage.NewDB()

	log.Println("Starting server...")

	err = http.ListenAndServe(os.Getenv("SERVER_PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
