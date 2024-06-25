package api

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func Run() {
	mux := http.NewServeMux()

	port := os.Getenv("PORT")

	// Account routes
	mux.HandleFunc("POST /Account/Create", handlerCreateAccount)
	mux.HandleFunc("PUT /Account/Update", onlyAdmin(handlerUpdateAccount))
	mux.HandleFunc("DELETE /Account/Delete/{id}", onlyAdmin(handlerDeleteAccount))
	mux.HandleFunc("GET /Account/GetAccounts", onlyAdmin(handlerGetAccounts))
	mux.HandleFunc("GET /Account/GetByID/{id}", onlyAdmin(handlerGetAccountByID))
	mux.HandleFunc("GET /Account/GetBySurname/{surname}", onlyAdmin(handlerGetAccountBySurname))
	mux.HandleFunc("POST /Account/Logging", handlerLogging)

	// CORS config
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := corsOptions.Handler(mux)

	log.Println("Running server...")

	log.Fatal(http.ListenAndServe(port, handler))
}
