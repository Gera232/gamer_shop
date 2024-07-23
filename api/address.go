package api

import (
	"back-end/storage"
	"back-end/types"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func handlerCreateAddress(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	addrs := types.Address{}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = json.Unmarshal(reqBody, &addrs)
	if err != nil {
		response := newResponse("Error", errUnmarshalFields.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = storage.CreateAddress(&addrs)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("Message", "successful creation", nil)
	responseJSON(w, http.StatusOK, response)
}

func handlerDeleteAddress(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	id := r.PathValue("id")

	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Para no hacerlo mas complejo no cheaqueare si existe o no.

	err = storage.DeleteAddress(uint32(id64))
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("Message", "successful removal", nil)
	responseJSON(w, http.StatusOK, response)
}

func handlerGetAddresses(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	id := r.PathValue("id")

	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	addrses, err := storage.GetAddresses(uint32(id64))
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", addrses)
	responseJSON(w, http.StatusOK, response)
}
