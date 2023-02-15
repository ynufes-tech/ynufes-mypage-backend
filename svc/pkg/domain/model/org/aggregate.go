package org

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	Org struct {
		ID     ID
		Event  event.Event
		Name   string
		Users  []user.ID
		IsOpen bool
	}
	ID util.ID
)
