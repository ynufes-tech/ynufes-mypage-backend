package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type Event interface {
	Create(context.Context, *event.Event) error
	UpdateName(context.Context, id.EventID, string) error
}
