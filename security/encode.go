package security

import (
	"encoding/base64"
)

func Encode(pass string) string {
	e := base64.StdEncoding.EncodeToString([]byte(pass))
	return e
}

func Decode(pass string) (string, error) {
	d, err := base64.StdEncoding.DecodeString(pass)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
