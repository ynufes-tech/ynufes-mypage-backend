package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/event"
)

type (
	Event struct {
		collection *firestore.CollectionRef
	}
)

func NewEvent(c *firestore.Client) Event {
	return Event{
		collection: c.Collection("events"),
	}
}

func (eve Event) Create(ctx context.Context, model *event.Event) error {
	e := &entity.Event{
		Name: model.Name,
	}
	_, err := eve.collection.Doc(model.ID.ExportID()).
		Create(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (eve Event) UpdateAll(ctx context.Context, model *event.Event) error {
	_, err := eve.collection.Doc(model.ID.ExportID()).
		Set(ctx, model)
	return err
}

func (eve Event) Delete(ctx context.Context, model *event.Event) error {
	_, err := eve.collection.Doc(model.ID.ExportID()).
		Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
