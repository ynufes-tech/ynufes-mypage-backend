package user

import (
	"time"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
	"ynufes-mypage-backend/svc/pkg/exception"
)

type (
	Agent struct {
		Roles []Role
	}
	Role struct {
		ID          RoleID
		Level       RoleLevel
		GrantedTime time.Time
	}
	RoleID    util.ID
	RoleLevel int
)

const (
	Viewer  RoleLevel = 1
	Editor  RoleLevel = 2
	Manager RoleLevel = 3
)

func NewRoleLevel(level int) (RoleLevel, error) {
	switch RoleLevel(level) {
	case Viewer:
		return Viewer, nil
	case Editor:
		return Editor, nil
	case Manager:
		return Manager, nil
	default:
		return 0, exception.ErrInvalidRoleLevel
	}
}
