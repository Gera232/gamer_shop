package handler

import (
	"back-end/database"
	model "back-end/model/account"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	account := model.Account{}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = json.Unmarshal(reqBody, &account)
	if err != nil {
		response := newResponse("Error", errUnmarshalFields.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	v := database.ExistAccount(account.Surname)
	log.Println(v)
	if v {
		response := newResponse("Error", errExistAccount.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = database.CreateAccount(&account)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("Message", "Account created", nil)
	responseJSON(w, http.StatusCreated, response)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {}

func GetAccounts(w http.ResponseWriter, r *http.Request) {}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	v := r.PathValue("id")

	v64, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	id := uint32(v64)

	account, err := database.GetAccountByID(id)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", account)
	responseJSON(w, http.StatusOK, response)
}

func GetAccountBySurname(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	surname := r.PathValue("surname")

	account, err := database.GetAccountBySurname(surname)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", account)
	responseJSON(w, http.StatusOK, response)
}

func Logging(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	account := model.Account{}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = json.Unmarshal(reqBody, &account)
	if err != nil {
		response := newResponse("Error", errUnmarshalFields.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	v, err := database.Logging(account.Surname, account.Password)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if !v {
		response := newResponse("Error", errAuthenticator.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
}
