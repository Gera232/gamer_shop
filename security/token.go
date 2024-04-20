package security

import (
	model "back-end/model/account"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name  string     `json:"name"`
	Admin model.Rols `json:"admin"`
	jwt.RegisteredClaims
}

func CreateToken(surname string, rol model.Rols) (string, error) {
	claims := &JwtCustomClaims{
		surname,
		rol,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte("Warrior"))
	if err != nil {
		return "", err
	}

	return token, nil
}
