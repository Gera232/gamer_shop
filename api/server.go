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

	// Account
	mux.HandleFunc("POST /CreateAccount", createAccount)
	mux.HandleFunc("PUT /UpdateAccount", onlyAdmin(updateAccount))
	mux.HandleFunc("DELETE /DeleteAccount/{id}", onlyAdmin(deleteAccount))
	mux.HandleFunc("GET /GetAccounts", onlyAdmin(getAccounts))
	mux.HandleFunc("GET /GetAccountByID/{id}", onlyAdmin(getAccountByID))
	mux.HandleFunc("GET /GetAccountBySurname/{surname}", onlyAdmin(getAccountBySurname))
	mux.HandleFunc("POST /logging", logging)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	server := corsOptions.Handler(mux)

	log.Println("Running server...")

	log.Fatal(http.ListenAndServe(port, server))
}
