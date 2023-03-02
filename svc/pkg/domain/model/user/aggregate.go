package user

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	User struct {
		ID     id.UserID
		Status Status
		Detail Detail
		Line   Line
		Admin  Admin
		Agent  Agent
	}
	Status int
)

const (
	// StatusNew indicates that user is newly created and hasn't finished its basic registration.
	StatusNew Status = 1
	// StatusRegistered indicates that user has finished its basic registration.
	StatusRegistered Status = 2
)

func (u User) IsValid() bool {
	return u.ID.HasValue() && u.Line.LineServiceID != ""
}
