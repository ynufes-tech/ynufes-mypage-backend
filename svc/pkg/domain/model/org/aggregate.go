package org

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	Org struct {
		ID     id.OrgID
		Event  event.Event
		Name   string
		IsOpen bool
	}
)
