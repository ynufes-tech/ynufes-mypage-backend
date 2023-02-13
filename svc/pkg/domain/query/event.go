package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
)

type Event interface {
	GetByID(ctx context.Context, id event.ID) (*event.Event, error)
}
