package api

import (
	model "back-end/model/account"
	"back-end/security"
	"back-end/storage"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func createAccount(w http.ResponseWriter, r *http.Request) {
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

	existAccount := storage.ExistAccount(account.Surname)
	if existAccount {
		response := newResponse("Error", errExistAccount.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = storage.CreateAccount(&account)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("Message", "Account created", nil)
	responseJSON(w, http.StatusCreated, response)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {}

func deleteAccount(w http.ResponseWriter, r *http.Request) {}

func getAccounts(w http.ResponseWriter, r *http.Request) {}

func getAccountByID(w http.ResponseWriter, r *http.Request) {
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

	account, err := storage.GetAccountByID(uint32(id64))
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", account)
	responseJSON(w, http.StatusOK, response)
}

func getAccountBySurname(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	surname := r.PathValue("surname")

	account, err := storage.GetAccountBySurname(surname)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", account)
	responseJSON(w, http.StatusOK, response)
}

func logging(w http.ResponseWriter, r *http.Request) {
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
	role, auhtenticator, err := storage.Logging(account.Surname, account.Password)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if !auhtenticator {
		response := newResponse("Error", errAuthenticator.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	token, err := security.CreateToken(role)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", token)
	responseJSON(w, http.StatusOK, response)
}
