package main

import (
	"back-end/database"
	"back-end/route"
	"net/http"
)

func main() {
	route.SetupRoutes()
	database.NewDB()

	http.ListenAndServe(":80", nil)
}
