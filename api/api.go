package api

import (
	"errors"
	"sync"
)

var (
	lock = sync.Mutex{}

	// errors
	errUnmarshalFields = errors.New("the data type of some field is wrong")
	errInternalServer  = errors.New("internal server error")
	errExistAccount    = errors.New("this account already exist")
	errAuthenticator   = errors.New("the surname or the password are wrong")
	errUnauthorized    = errors.New("unauthorized")
	errAccountNotExist = errors.New("this account not exist")
)
