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
		Users  []id.UserID
		IsOpen bool
	}
)

func (o Org) IsGranted(userID id.UserID) bool {
	for _, u := range o.Users {
		if u == userID {
			return true
		}
	}
	return false
}
