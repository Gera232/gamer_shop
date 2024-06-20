package api

import (
	"encoding/json"
	"errors"
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

type responseStruck struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newResponse(messageType string, message string, data interface{}) responseStruck {
	return responseStruck{
		messageType,
		message,
		data,
	}
}

func responseJSON(w http.ResponseWriter, statusCode int, resp responseStruck) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
