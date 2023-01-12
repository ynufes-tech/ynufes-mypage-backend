package org

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	Org struct {
		ID      ID
		EventID event.ID
		Name    string
		Members []user.ID
	}
	ID util.ID
)
