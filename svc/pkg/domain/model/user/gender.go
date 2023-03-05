package user

import "errors"

type Gender int

const (
	GenderUnknown      Gender = 0
	GenderMan          Gender = 1
	GenderWoman        Gender = 2
	GenderNotSpecified Gender = 3
)

func NewGender(gender int) (Gender, error) {
	switch Gender(gender) {
	case GenderMan:
		return GenderMan, nil
	case GenderWoman:
		return GenderWoman, nil
	case GenderNotSpecified:
		return GenderNotSpecified, nil
	case GenderUnknown:
		return GenderUnknown, nil
	default:
		return -1, errors.New("GENDER VALUE IS INVALID")
	}
}
