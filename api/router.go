package api

import (
	"net/http"
)

func SetupRoutes() {
	// Account
	http.HandleFunc("POST /CreateAccount", createAccount)
	http.HandleFunc("PUT /UpdateAccount/{id}", updateAccount)
	http.HandleFunc("DELETE /DeleteAccount/{id}", deleteAccount)
	http.HandleFunc("GET /GetAccounts", getAccounts)
	http.HandleFunc("GET /GetAccountByID/{id}", getAccountByID)
	http.HandleFunc("GET /GetAccountBySurname/{surname}", getAccountBySurname)
	http.HandleFunc("GET /logging", logging)
}
