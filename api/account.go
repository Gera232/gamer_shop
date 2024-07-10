package api

import (
	"back-end/security"
	"back-end/storage"
	"back-end/types"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func handlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	acc := types.Account{}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = json.Unmarshal(reqBody, &acc)
	if err != nil {
		response := newResponse("Error", errUnmarshalFields.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	existAcc := storage.ExistAccountSurname(acc.Surname)
	if existAcc {
		response := newResponse("Error", errExistAccount.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = storage.CreateAccount(&acc)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("Message", "successful creation", nil)
	responseJSON(w, http.StatusOK, response)
}

func handlerUpdateAccount(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	acc := types.Account{}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = json.Unmarshal(reqBody, &acc)
	if err != nil {
		response := newResponse("Error", errUnmarshalFields.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	existAcc := storage.ExistAccountSurname(acc.Surname)
	if !existAcc {
		response := newResponse("Error", errAccountNotExist.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = storage.UpdateAccount(&acc)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("Message", "successful modify", nil)
	responseJSON(w, http.StatusOK, response)
}

func handlerDeleteAccount(w http.ResponseWriter, r *http.Request) {
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

	existAcc := storage.ExistAccountID(uint32(id64))
	if !existAcc {
		response := newResponse("Error", errAccountNotExist.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = storage.DeleteAccount(uint32(id64))
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("Message", "successful removal", nil)
	responseJSON(w, http.StatusOK, response)
}

func handlerGetAccounts(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	accs, err := storage.GetAccounts()
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", accs)
	responseJSON(w, http.StatusOK, response)
}

func handlerGetAccountByID(w http.ResponseWriter, r *http.Request) {
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

	existAcc := storage.ExistAccountID(uint32(id64))
	if !existAcc {
		response := newResponse("Error", errAccountNotExist.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
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

func handlerGetAccountBySurname(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	surname := r.PathValue("surname")

	existAcc := storage.ExistAccountSurname(surname)
	if !existAcc {
		response := newResponse("Error", errAccountNotExist.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	acc, err := storage.GetAccountBySurname(surname)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", acc)
	responseJSON(w, http.StatusOK, response)
}

func handlerLogging(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	acc := types.Account{}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = json.Unmarshal(reqBody, &acc)
	if err != nil {
		response := newResponse("Error", errUnmarshalFields.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	role, id, auht, err := storage.Logging(acc.Surname, acc.Password)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if !auht {
		response := newResponse("Error", errAuthenticator.Error(), nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	token, err := security.CreateToken(role, id)
	if err != nil {
		log.Println(err)
		response := newResponse("Error", errInternalServer.Error(), nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("", "", token)
	responseJSON(w, http.StatusOK, response)
}
