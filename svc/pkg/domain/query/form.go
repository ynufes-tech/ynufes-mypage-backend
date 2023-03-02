package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type Form interface {
	GetByID(ctx context.Context, id id.FormID) (*form.Form, error)
	ListByEventID(ctx context.Context, eventID id.EventID) ([]form.Form, error)
}
