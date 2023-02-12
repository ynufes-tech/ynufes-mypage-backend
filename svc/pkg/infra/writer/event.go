package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/event"
)

type (
	Event struct {
		collection *firestore.CollectionRef
	}
)

func NewEvent(c *firestore.Client) Event {
	return Event{
		collection: c.Collection(entity.EventCollectionName),
	}
}

func (eve Event) Create(ctx context.Context, model event.Event) error {
	if !model.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}
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
