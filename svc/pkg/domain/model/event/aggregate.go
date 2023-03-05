package event

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	Event struct {
		ID   id.EventID
		Name string
	}
)
