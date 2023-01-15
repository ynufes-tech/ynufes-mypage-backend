package query

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
)

type Event interface {
	GetByID(id event.ID) (*event.Event, error)
}
