package api

import (
	"errors"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/rs/cors"
)

var (
	lock = sync.Mutex{}

	// errors
	errUnmarshalFields = errors.New("the data type of some field is wrong")
	errInternalServer  = errors.New("internal server error")
	errExistAccount    = errors.New("this account already exist")
	errAuthenticator   = errors.New("the surname or the password are wrong")
	errUnauthorized    = errors.New("unauthorized")
	errAccountNotExist = errors.New("this account not exist")
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

	// Address routes
	mux.HandleFunc("POST /Address/Create", handlerCreateAddress)
	mux.HandleFunc("DELETE /Address/Delete/{id}", handlerDeleteAddress)
	mux.HandleFunc("GET /Address/GetAddresses/{id}", handlerGetAddresses)

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
