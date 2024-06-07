package main

import (
	"backend/api"
	"backend/storage"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	time.Sleep(time.Second * 5)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	storage.NewDB()

	api.SetupRoutes()

	log.Println("Powering up server...")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
