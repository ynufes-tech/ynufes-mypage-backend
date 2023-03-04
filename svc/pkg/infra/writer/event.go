package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/event"
	orgEntity "ynufes-mypage-backend/svc/pkg/infra/entity/org"
)

type (
	Event struct {
		eventRef *db.Ref
		orgRef   *db.Ref
	}
)

func NewEvent(f *firebase.Firebase) Event {
	return Event{
		eventRef: f.Client(entity.EventRootName),
		orgRef:   f.Client(orgEntity.OrgRootName),
	}
}

func (eve Event) Create(ctx context.Context, model *event.Event) error {
	if model.ID != nil && model.ID.HasValue() {
		return exception.ErrIDAlreadyAssigned
	}
	newID := identity.IssueID()
	e := entity.Event{
		Name: model.Name,
	}
	err := eve.eventRef.Child(newID.ExportID()).Set(ctx, e)
	if err != nil {
		return err
	}
	model.ID = newID
	return nil
}

func (eve Event) Set(ctx context.Context, model event.Event) error {
	if !model.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	e := entity.Event{
		Name: model.Name,
	}
	err := eve.eventRef.Child(model.ID.ExportID()).Set(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

// UpdateName updates name in EventRoot.
// event_name in Orgs should be updated by GoogleCloudFunctions or as concurrent process.
// TODO: implement GoogleCloudFunctions or concurrent process.
func (eve Event) UpdateName(ctx context.Context, tid id.EventID, name string) error {
	err := eve.eventRef.Child(tid.ExportID()).
		Transaction(
			ctx,
			func(t db.TransactionNode) (interface{}, error) {
				var target entity.Event
				if err := t.Unmarshal(&target); err != nil {
					return nil, err
				}
				target.Name = name
				return target, nil
			},
		)
	if err != nil {
		return fmt.Errorf("failed to update name of Event: %w", err)
	}
	return nil
}
