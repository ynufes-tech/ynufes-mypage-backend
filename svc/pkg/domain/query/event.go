package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type Event interface {
	GetByID(ctx context.Context, id id.EventID) (*event.Event, error)
}
