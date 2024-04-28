package security

import (
	"encoding/base64"
)

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
