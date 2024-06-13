package api

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func Run(port string) {
	mux := http.NewServeMux()

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

	log.Printf("Starting server on port%v", port)

	log.Fatal(http.ListenAndServe(port, server))
}
