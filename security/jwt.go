package security

import (
	model "back-end/model/account"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Surname string     `json:"surname"`
	Role    model.Rols `json:"role"`
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	CodifiedToken, err := token.SignedString([]byte("Pava electrica"))
	if err != nil {
		return "", err
	}

	return CodifiedToken, nil
}
