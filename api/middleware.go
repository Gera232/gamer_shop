package api

import (
	"back-end/security"
	"back-end/storage"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func onlyAdmin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-jwt-token")

		validatedToken, err := security.ValidateJWT(token)
		if err != nil {
			log.Println(err)
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		if !validatedToken.Valid {
			log.Println(err)
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		claims := validatedToken.Claims.(jwt.MapClaims)

		id, ok := claims["id"].(float64)
		if !ok {
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		if !storage.ExistAccountID(uint32(id)) {
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		if claims["role"] != "admin" {
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		f(w, r)
	}
}
