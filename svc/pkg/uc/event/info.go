package event

import (
	"errors"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type InfoUseCaseInput struct {
	ID event.ID
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

func (uc InfoUseCase) Do(input InfoUseCaseInput) (InfoUseCaseOutput, error) {
	eventD, err := uc.eventQ.GetByID(input.ID)
	if err != nil {
		return InfoUseCaseOutput{}, errors.New("event not found")
	}

	return InfoUseCaseOutput{
		Event: *eventD,
	}, nil
}
