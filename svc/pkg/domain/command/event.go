package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
)

type Event interface {
	Create(context.Context, *event.Event) error
	UpdateName(context.Context, *event.Event) error
}
