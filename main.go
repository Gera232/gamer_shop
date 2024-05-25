package main

import (
	"back-end/api"
	"back-end/storage"
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

	log.Println("all good")

	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
