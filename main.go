package main

import (
	"backend/api"
	"backend/storage"
	"log"
	"net/http"
	"time"
)

func main() {
	time.Sleep(time.Second * 5)

	storage.NewDB()

	api.SetupRoutes()

	log.Println("Powering up server...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
