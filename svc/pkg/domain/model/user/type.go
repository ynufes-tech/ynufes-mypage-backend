package user

import "errors"

// Type represents the type of user.
// This field is currently not used, but it is expected to be used in the future.
type Type int

const (
	TypeNormal Type = 1
)

func NewType(t int) (Type, error) {
	switch Type(t) {
	case TypeNormal:
		return TypeNormal, nil
	default:
		return -1, errors.New("USER TYPE VALUE IS INVALID")
	}
}
