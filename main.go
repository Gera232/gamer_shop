package main

import (
	"backend/api"
	"backend/storage"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("./database.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	storage.NewDB()

	err = storage.Migrate(file)
	if err != nil {
		log.Fatal(err)
	}

	api.SetupRoutes()

	port := os.Getenv("PORT")

	log.Printf("Starting server on port: %v", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
