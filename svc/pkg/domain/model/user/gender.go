package user

import "errors"

type Gender int

const (
	GenderMan          = 1
	GenderWoman        = 2
	GenderNotSpecified = 0
)

func NewGender(gender int) (Gender, error) {
	switch Gender(gender) {
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
