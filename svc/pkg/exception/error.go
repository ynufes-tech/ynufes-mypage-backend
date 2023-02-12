package exception

import "errors"

var (
	ErrorInvalidHeader = errors.New("INVALID Authorization Header")
	ErrInvalidJWT      = errors.New("INVALID JWT")
	ErrIDNotAssigned   = errors.New("ID NOT ASSIGNED")
)
