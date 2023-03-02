package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/org"
)

type Org struct {
	collection *firestore.CollectionRef
}

func NewOrg(c *firestore.Client) Org {
	return Org{
		collection: c.Collection(entity.OrgCollectionName),
	}
}

func (w Org) Create(ctx context.Context, o org.Org) error {
	if !o.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	usersE := make([]int64, len(o.Users))
	for i := range o.Users {
		usersE[i] = o.Users[i].GetValue()
	}
	e := entity.Org{
		EventID:   o.Event.ID.GetValue(),
		EventName: o.Event.Name,
		Name:      o.Name,
		Users:     usersE,
		IsOpen:    o.IsOpen,
	}
	if _, err := w.collection.Doc(o.ID.ExportID()).Create(ctx, e); err != nil {
		log.Printf("Failed to create org: %v", err)
		return fmt.Errorf("failed to create org: %w", err)
	}
	return nil
}

func (w Org) Set(ctx context.Context, o org.Org) error {
	usersE := make([]int64, len(o.Users))
	e := entity.Org{
		EventID:   o.Event.ID.GetValue(),
		EventName: o.Event.Name,
		Name:      o.Name,
		Users:     usersE,
		IsOpen:    o.IsOpen,
	}
	for i := range o.Users {
		usersE[i] = o.Users[i].GetValue()
	}
	_, err := w.collection.Doc(o.ID.ExportID()).Set(ctx, e)
	if err != nil {
		log.Printf("Failed to update org: %v", err)
		return fmt.Errorf("failed to update org: %w", err)
	}
	return nil
}

func (w Org) UpdateUsers(ctx context.Context, o org.Org) error {
	usersE := make([]int64, len(o.Users))
	for i := range o.Users {
		usersE[i] = o.Users[i].GetValue()
	}
	_, err := w.collection.Doc(o.ID.ExportID()).Update(ctx,
		[]firestore.Update{
			{Path: "user_ids", Value: usersE},
		})
	if err != nil {
		log.Printf("Failed to update org users: %v", err)
		return fmt.Errorf("failed to update org users: %w", err)
	}
	return nil
}

func (w Org) UpdateIsOpen(ctx context.Context, o org.Org) error {
	_, err := w.collection.Doc(o.ID.ExportID()).Update(ctx,
		[]firestore.Update{
			{Path: "is_open", Value: o.IsOpen},
		})
	if err != nil {
		log.Printf("Failed to update org is_open: %v", err)
		return fmt.Errorf("failed to update org is_open: %w", err)
	}
	return nil
}
