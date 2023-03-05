package event

import (
	"context"
	"errors"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type InfoUseCaseInput struct {
	ID id.EventID
}

type InfoUseCaseOutput struct {
	Event event.Event
}

type InfoUseCase struct {
	eventQ query.Event
}

func NewInfoUseCase(rgst registry.Registry) InfoUseCase {
	return InfoUseCase{
		// TODO: IMPLEMENT ME
	}
}

func (uc InfoUseCase) Do(ctx context.Context, input InfoUseCaseInput) (InfoUseCaseOutput, error) {
	eventD, err := uc.eventQ.GetByID(ctx, input.ID)
	if err != nil {
		return InfoUseCaseOutput{}, errors.New("event not found")
	}

	return InfoUseCaseOutput{
		Event: *eventD,
	}, nil
}
