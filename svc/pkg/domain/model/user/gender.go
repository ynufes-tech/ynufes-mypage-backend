package user

import "errors"

type Gender int

const (
	GenderMan          = 0
	GenderWoman        = 1
	GenderNotSpecified = 2
)

func NewGender(gender int) (Gender, error) {
	switch gender {
	case GenderMan:
		return GenderMan, nil
	case GenderWoman:
		return GenderWoman, nil
	case GenderNotSpecified:
		return GenderNotSpecified, nil
	default:
		return -1, errors.New("GENDER VALUE IS INVALID")
	}
}
