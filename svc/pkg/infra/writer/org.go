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
	memberE := make([]int64, len(org.Members))
	for i := range org.Members {
		memberE[i] = org.Members[i].GetValue()
	}
	e := entity.Org{
		EventID:   org.Event.ID.GetValue(),
		EventName: org.Event.Name,
		Name:      org.Name,
		Members:   memberE,
		IsOpen:    org.IsOpen,
	}
	if _, err := o.collection.Doc(org.ID.ExportID()).Create(ctx, e); err != nil {
		log.Printf("Failed to create org: %v", err)
		return fmt.Errorf("failed to create org: %w", err)
	}
	return nil
}

func (o Org) Set(ctx context.Context, org org.Org) error {
	memberE := make([]int64, len(org.Members))
	e := entity.Org{
		EventID:   org.Event.ID.GetValue(),
		EventName: org.Event.Name,
		Name:      org.Name,
		Members:   memberE,
		IsOpen:    org.IsOpen,
	}
	for i := range org.Members {
		memberE[i] = org.Members[i].GetValue()
	}
	_, err := o.collection.Doc(org.ID.ExportID()).Set(ctx, e)
	if err != nil {
		log.Printf("Failed to update org: %v", err)
		return fmt.Errorf("failed to update org: %w", err)
	}
	return nil
}

func (o Org) UpdateMembers(ctx context.Context, org org.Org) error {
	memberE := make([]int64, len(org.Members))
	for i := range org.Members {
		memberE[i] = org.Members[i].GetValue()
	}
	_, err := o.collection.Doc(org.ID.ExportID()).Update(ctx,
		[]firestore.Update{
			{Path: "member_ids", Value: memberE},
		})
	if err != nil {
		log.Printf("Failed to update org members: %v", err)
		return fmt.Errorf("failed to update org members: %w", err)
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
