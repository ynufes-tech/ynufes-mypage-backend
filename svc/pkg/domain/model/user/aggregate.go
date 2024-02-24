package user

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	User struct {
		ID     id.UserID
		Detail Detail
		Agent  Agent
	}
)

func (u User) IsValid() bool {
	return u.ID != nil && u.ID.HasValue()
}
