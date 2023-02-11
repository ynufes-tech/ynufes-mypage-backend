package org

import (
	"errors"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

var ErrAlreadyHasID = errors.New("org struct already has id")

type (
	Org struct {
		ID      ID
		Event   event.Event
		Name    string
		Members []user.ID
		IsOpen  bool
	}
	ID util.ID
)

func (o *Org) AssignID(id ID) error {
	if id.HasValue() {
		return ErrAlreadyHasID
	}
	o.ID = id
	return nil
}
