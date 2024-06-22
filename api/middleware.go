package api

import (
	"api/security"
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
		if claims["role"] != "admin" {
			log.Println(claims["role"])
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		f(w, r)
	}
}
