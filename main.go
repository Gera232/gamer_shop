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
	port := os.Getenv("PORT")

	time.Sleep(time.Second * 5)

	storage.NewDB()

	api.SetupRoutes()

	log.Printf("Starting server on port %v", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
