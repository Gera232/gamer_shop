package main

import (
	"backend/api"
	"backend/storage"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	time.Sleep(time.Second * 5)

	storage.NewDB()

	api.SetupRoutes()

	port := os.Getenv("PORT")

	log.Printf("Starting server on port: %v", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
