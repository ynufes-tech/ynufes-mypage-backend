package user

import (
	"time"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
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
	Editor            = 2
	Manager           = 3
)
