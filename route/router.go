package route

import (
	"back-end/handler"
	"net/http"
)

func SetupRoutes() {
	// Account
	http.HandleFunc("POST /CreateAccount", handler.CreateAccount)
	http.HandleFunc("PUT /UpdateAccount/{id}", handler.UpdateAccount)
	http.HandleFunc("DELETE /DeleteAccount/{id}", handler.DeleteAccount)
	http.HandleFunc("GET /GetAccounts", handler.GetAccounts)
	http.HandleFunc("GET /GetOneAccount/{id}", handler.GetOneAccount)
}
