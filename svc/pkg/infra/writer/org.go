package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
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

func (o Org) Create(ctx context.Context, org *org.Org) error {
	memberE := make([]int64, len(org.Members))
	for i := range org.Members {
		memberE[i] = org.Members[i].GetValue()
	}
	id := identity.IssueID()
	e := entity.Org{
		EventID:   org.Event.ID.GetValue(),
		EventName: org.Event.Name,
		Name:      org.Name,
		Members:   memberE,
		IsOpen:    org.IsOpen,
	}
	_, err := o.collection.Doc(id.ExportID()).Create(ctx, e)
	if err != nil {
		log.Printf("Failed to create org: %v", err)
		return fmt.Errorf("failed to create org: %w", err)
	}
	org.AssignID(org.ID)
	return nil
}
