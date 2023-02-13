package event

import (
	"context"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type CreateUseCase struct {
	eventC command.Event
}

type CreateInput struct {
	Ctx       context.Context
	EventName string
}

type CreateOutput struct {
	Event event.Event
}

func NewCreate(rgst registry.Registry) *CreateUseCase {
	return &CreateUseCase{
		eventC: rgst.Repository().NewEventCommand(),
	}
}

func (c CreateUseCase) Do(ipt CreateInput) (*CreateOutput, error) {
	e := event.Event{
		Name: ipt.EventName,
		ID:   identity.IssueID(),
	}
	if err := c.eventC.Create(ipt.Ctx, e); err != nil {
		return nil, err
	}
	return &CreateOutput{
		Event: e,
	}, nil
}
