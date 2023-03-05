package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"log"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/org"
)

type Org struct {
	ref *db.Ref
}

func NewOrg(f *firebase.Firebase) Org {
	return Org{
		ref: f.Client(entity.OrgRootName),
	}
}

// Create new ID will be generated and assigned,
// do not assign ID to the argument.
func (w Org) Create(ctx context.Context, o *org.Org) error {
	if o.ID != nil && o.ID.HasValue() {
		return exception.ErrIDAlreadyAssigned
	}
	newID := identity.IssueID()
	e := entity.Org{
		EventID:   o.Event.ID.ExportID(),
		EventName: o.Event.Name,
		Name:      o.Name,
		IsOpen:    o.IsOpen,
	}
	if err := w.ref.Child(newID.ExportID()).
		Set(ctx, e); err != nil {
		log.Printf("Failed to create org: %v", err)
		return fmt.Errorf("failed to create org: %w", err)
	}
	o.ID = newID
	return nil
}

func (w Org) Set(ctx context.Context, o org.Org) error {
	if !o.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	e := entity.Org{
		EventID:   o.Event.ID.ExportID(),
		EventName: o.Event.Name,
		Name:      o.Name,
		IsOpen:    o.IsOpen,
	}
	if err := w.ref.Child(o.ID.ExportID()).
		Set(ctx, e); err != nil {
		log.Printf("Failed to update org: %v", err)
		return fmt.Errorf("failed to update org: %w", err)
	}
	return nil
}

func (w Org) UpdateIsOpen(ctx context.Context, tID id.OrgID, isOpen bool) error {
	if err := w.ref.Child(tID.ExportID()).Child("is_open").
		Update(ctx, map[string]interface{}{
			"is_open": isOpen,
		}); err != nil {
		log.Printf("Failed to update org is_open: %v", err)
		return fmt.Errorf("failed to update org is_open: %w", err)
	}
	return nil
}
