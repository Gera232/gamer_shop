package security

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(role string) (string, error) {
	claims := &jwt.MapClaims{
		"role":      role,
		"expiresAT": 15000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	CodifiedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return CodifiedToken, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
}
