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

func (o Org) Create(ctx context.Context, org org.Org) error {
	if !org.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	usersE := make([]int64, len(org.Users))
	for i := range org.Users {
		usersE[i] = org.Users[i].GetValue()
	}
	e := entity.Org{
		EventID:   org.Event.ID.GetValue(),
		EventName: org.Event.Name,
		Name:      org.Name,
		Users:     usersE,
		IsOpen:    org.IsOpen,
	}
	if _, err := o.collection.Doc(org.ID.ExportID()).Create(ctx, e); err != nil {
		log.Printf("Failed to create org: %v", err)
		return fmt.Errorf("failed to create org: %w", err)
	}
	return nil
}

func (o Org) Set(ctx context.Context, org org.Org) error {
	usersE := make([]int64, len(org.Users))
	e := entity.Org{
		EventID:   org.Event.ID.GetValue(),
		EventName: org.Event.Name,
		Name:      org.Name,
		Users:     usersE,
		IsOpen:    org.IsOpen,
	}
	for i := range org.Users {
		usersE[i] = org.Users[i].GetValue()
	}
	_, err := o.collection.Doc(org.ID.ExportID()).Set(ctx, e)
	if err != nil {
		log.Printf("Failed to update org: %v", err)
		return fmt.Errorf("failed to update org: %w", err)
	}
	return nil
}

func (o Org) UpdateUsers(ctx context.Context, org org.Org) error {
	usersE := make([]int64, len(org.Users))
	for i := range org.Users {
		usersE[i] = org.Users[i].GetValue()
	}
	_, err := o.collection.Doc(org.ID.ExportID()).Update(ctx,
		[]firestore.Update{
			{Path: "user_ids", Value: usersE},
		})
	if err != nil {
		log.Printf("Failed to update org users: %v", err)
		return fmt.Errorf("failed to update org users: %w", err)
	}
	return nil
}

func (o Org) UpdateIsOpen(ctx context.Context, org org.Org) error {
	_, err := o.collection.Doc(org.ID.ExportID()).Update(ctx,
		[]firestore.Update{
			{Path: "is_open", Value: org.IsOpen},
		})
	if err != nil {
		log.Printf("Failed to update org is_open: %v", err)
		return fmt.Errorf("failed to update org is_open: %w", err)
	}
	return nil
}
