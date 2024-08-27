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
	mux.HandleFunc("POST /api/account/create", handlerCreateAccount)
	mux.HandleFunc("PUT /api/account/update", onlyAdmin(handlerUpdateAccount))
	mux.HandleFunc("DELETE /api/account/delete/{id}", onlyAdmin(handlerDeleteAccount))
	mux.HandleFunc("GET /api/account/getAccounts", onlyAdmin(handlerGetAccounts))
	mux.HandleFunc("GET /api/account/getByID/{id}", onlyAdmin(handlerGetAccountByID))
	mux.HandleFunc("GET /api/account/getBySurname/{surname}", onlyAdmin(handlerGetAccountBySurname))
	mux.HandleFunc("POST /api/account/login", handlerLogin)

	// CORS config
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsOptions.Handler(mux)

	log.Println("Running server...")

	log.Fatal(http.ListenAndServe(port, handler))
}
