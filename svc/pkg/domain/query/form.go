package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
)

type Form interface {
	GetByID(ctx context.Context, id form.ID) (*form.Form, error)
	ListByEventID(ctx context.Context, eventID event.ID) ([]form.Form, error)
}
