package event

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	Event struct {
		ID   ID
		Name string
	}
	ID util.ID
)
