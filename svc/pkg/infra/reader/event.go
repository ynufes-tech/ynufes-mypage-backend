package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/event"
)

type Event struct {
	ref *db.Ref
}

func NewEvent(c *firebase.Firebase) Event {
	return Event{
		ref: c.Client(entity.EventRootName),
	}
}

func (e Event) GetByID(ctx context.Context, id id.EventID) (*event.Event, error) {
	if !id.HasValue() {
		return nil, exception.ErrIDNotAssigned
	}
	var eventEntity entity.Event

	r, err := e.ref.OrderByKey().EqualTo(id.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	if len(r) > 1 {
		fmt.Printf("multiple event found with id: %s", id)
	}
	if err := r[0].Unmarshal(&eventEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event entity: %w", err)
	}
	eventEntity.ID = id
	model, err := eventEntity.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert event entity to model: %w", err)
	}
	return model, nil
}
