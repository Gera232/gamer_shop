package api

import (
	"net/http"
)

func SetupRoutes() {
	// Migrate
	http.HandleFunc("POST /Migrate", migrate)

	// Account
	http.HandleFunc("POST /CreateAccount", createAccount)
	http.HandleFunc("PUT /UpdateAccount", onlyAdmin(updateAccount))
	http.HandleFunc("DELETE /DeleteAccount/{id}", onlyAdmin(deleteAccount))
	http.HandleFunc("GET /GetAccounts", onlyAdmin(getAccounts))
	http.HandleFunc("GET /GetAccountByID/{id}", onlyAdmin(getAccountByID))
	http.HandleFunc("GET /GetAccountBySurname/{surname}", onlyAdmin(getAccountBySurname))
	http.HandleFunc("GET /logging", logging)
}
