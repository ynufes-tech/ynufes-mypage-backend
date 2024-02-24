package staff

import "ynufes-mypage-backend/svc/pkg/domain/model/id"

type Staff struct {
	id.UserID
	IsAdmin bool
}
