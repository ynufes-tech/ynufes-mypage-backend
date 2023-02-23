package form

import "errors"

type (
	FormWithStatus struct {
		Form
		Status Status
	}
	Status int
)

const (
	Accepted     Status = 1
	Submitted    Status = 2
	NotSubmitted Status = 3
)

func NewFormWithStatus(form Form, status Status) *FormWithStatus {
	return &FormWithStatus{
		Form:   form,
		Status: status,
	}
}

func NewStatus(status int) (Status, error) {
	switch Status(status) {
	case Accepted:
		return Accepted, nil
	case Submitted:
		return Submitted, nil
	case NotSubmitted:
		return NotSubmitted, nil
	default:
		return -1, errors.New("STATUS VALUE IS INVALID")
	}
}
