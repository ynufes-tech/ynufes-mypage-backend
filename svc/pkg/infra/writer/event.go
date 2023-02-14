package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/event"
	orgEntity "ynufes-mypage-backend/svc/pkg/infra/entity/org"
)

type (
	Event struct {
		client          *firestore.Client
		collectionEvent *firestore.CollectionRef
		collectionOrg   *firestore.CollectionRef
	}
)

func NewEvent(c *firestore.Client) Event {
	return Event{
		client:          c,
		collectionEvent: c.Collection(entity.EventCollectionName),
		collectionOrg:   c.Collection(orgEntity.OrgCollectionName),
	}
}

func (eve Event) Create(ctx context.Context, model event.Event) error {
	if !model.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	e := &entity.Event{
		Name: model.Name,
	}
	_, err := eve.collectionEvent.Doc(model.ID.ExportID()).
		Create(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

// UpdateName updates name of event, including event_name in Orgs and name in Events
func (eve Event) UpdateName(ctx context.Context, model event.Event) error {
	err := eve.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		orgsRef := eve.collectionOrg.Where("event_id", "==", model.ID.GetValue())
		orgs, err := tx.Documents(orgsRef).GetAll()
		if err != nil {
			return err
		}
		for _, org := range orgs {
			err := tx.Update(org.Ref, []firestore.Update{
				{Path: "event_name", Value: model.Name},
			})
			if err != nil {
				return err
			}
		}
		if err := tx.Update(eve.collectionEvent.Doc(model.ID.ExportID()), []firestore.Update{
			{Path: "name", Value: model.Name},
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}
	return nil
}

func (eve Event) Delete(ctx context.Context, model *event.Event) error {
	_, err := eve.collectionEvent.Doc(model.ID.ExportID()).
		Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
