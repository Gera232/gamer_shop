package security

import (
	"encoding/base64"
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

func Encode(pass string) string {
	encodePass := base64.StdEncoding.EncodeToString([]byte(pass))
	return encodePass
}

func Decode(pass string) (string, error) {
	decodePass, err := base64.StdEncoding.DecodeString(pass)
	if err != nil {
		return "", err
	}
	return string(decodePass), nil
}
