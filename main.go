package main

import (
	"backend/api"
	"backend/storage"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	api.SetupRoutes()

	storage.NewDB()

	log.Println("Powering up server...")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
