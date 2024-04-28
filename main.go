package main

import (
	"back-end/api"
	"back-end/storage"
	"log"
	"net/http"
)

func main() {
	api.SetupRoutes()

	storage.NewDB()

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
