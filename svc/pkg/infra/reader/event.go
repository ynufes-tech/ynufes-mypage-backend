package reader

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/event"
)

type Event struct {
	collection *firestore.CollectionRef
}

func NewEvent(c *firestore.Client) Event {
	return Event{
		collection: c.Collection("events"),
	}
}

func (e Event) GetByID(ctx context.Context, id event.ID) (model *event.Event, err error) {
	var eventEntity entity.Event
	snap, err := e.collection.Doc(id.ExportID()).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get event doc: %w", err)
	}
	err = snap.DataTo(&eventEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to decode event doc into entity: %w", err)
	}
	eventEntity.ID = id
	model, err = eventEntity.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert event entity to model: %w", err)
	}
	return model, nil
}
