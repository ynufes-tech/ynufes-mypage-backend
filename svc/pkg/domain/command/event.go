package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
)

type Event interface {
	Create(context.Context, event.Event) error
	UpdateAll(context.Context, *event.Event) error
	Delete(context.Context, *event.Event) error
}
