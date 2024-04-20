package main

import (
	"back-end/database"
	"back-end/route"
	"log"
	"net/http"
)

func main() {
	route.SetupRoutes()

	database.NewDB()

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
