package user

import "errors"

type Type int

const (
	TypeNormal Type = 0
)

func NewType(t int) (Type, error) {
	switch t {
	case int(TypeNormal):
		return TypeNormal, nil
	default:
		return -1, errors.New("USER TYPE VALUE IS INVALID")
	}
}
