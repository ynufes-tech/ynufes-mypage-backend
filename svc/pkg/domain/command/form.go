package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
)

type Form interface {
	Create(context.Context, event.ID, *form.Form) error
	Delete(context.Context, event.ID, *form.Form) error
	Grant(context.Context, event.ID, *form.Form) error
}
