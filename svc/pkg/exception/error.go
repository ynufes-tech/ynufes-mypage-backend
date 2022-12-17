package exception

import "errors"

var (
	ErrorInvalidHeader = errors.New("invalid Authorization Header")
)
