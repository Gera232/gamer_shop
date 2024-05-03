package api

import (
	"back-end/security"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func onlyAdmin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-jwt-token")

		token, err := security.ValidateJWT(tokenString)
		if err != nil {
			log.Println(err)
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			log.Println(err)
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if claims["role"] != "admin" {
			log.Println(claims["role"])
			response := newResponse("Error", errUnauthorized.Error(), nil)
			responseJSON(w, http.StatusUnauthorized, response)
			return
		}

		f(w, r)
	}
}
