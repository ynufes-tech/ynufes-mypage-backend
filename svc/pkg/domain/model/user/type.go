package user

import "errors"

type Type int

const (
	TypeNormal Type = 1
	TypeMember Type = 2
)

func NewType(t int) (Type, error) {
	switch Type(t) {
	case TypeNormal:
		return TypeNormal, nil
	case TypeMember:
		return TypeMember, nil
	default:
		return -1, errors.New("USER TYPE VALUE IS INVALID")
	}
}
