package api

import (
	"backend/storage"
	"errors"
	"log"
	"net/http"
	"sync"
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

func migrate(w http.ResponseWriter, r *http.Request) {
	err := storage.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	response := newResponse("Message", "successful migration", nil)
	responseJSON(w, http.StatusOK, response)
}
