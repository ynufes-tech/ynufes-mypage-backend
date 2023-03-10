package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/form"
)

type Form struct {
	ref *db.Ref
}

func NewForm(f *firebase.Firebase) *Form {
	return &Form{
		ref: f.Client(entity.FormRootName),
	}
}

func (f Form) Create(ctx context.Context, target *form.Form) error {
	if target.ID != nil && target.ID.HasValue() {
		return exception.ErrIDAlreadyAssigned
	}
	tid := identity.IssueID()

	var roles = make(map[string]bool, len(target.Roles))
	for i := 0; i < len(target.Roles); i++ {
		roles[target.Roles[i].ExportID()] = true
	}
	e := entity.NewForm(
		tid,
		target.EventID.ExportID(),
		target.Title,
		target.Summary,
		target.Description,
		target.Sections,
		roles,
		target.Deadline.UnixMilli(),
		target.IsOpen,
	)
	err := f.ref.Child(tid.ExportID()).Set(ctx, e)
	if err != nil {
		return err
	}
	target.ID = tid
	return nil
}

func (f Form) Set(ctx context.Context, target form.Form) error {
	if target.ID == nil || !target.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}

	var roles = make(map[string]bool, len(target.Roles))
	for i := 0; i < len(target.Roles); i++ {
		roles[target.Roles[i].ExportID()] = true
	}
	e := entity.NewForm(
		target.ID,
		target.EventID.ExportID(),
		target.Title,
		target.Summary,
		target.Description,
		target.Sections,
		roles,
		target.Deadline.UnixMilli(),
		target.IsOpen,
	)
	err := f.ref.Child(target.ID.ExportID()).Set(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (f Form) AddSectionOrder(ctx context.Context, fid id.FormID, sid id.SectionID, index float64) error {
	if fid == nil || !fid.HasValue() ||
		sid == nil || !sid.HasValue() {
		return exception.ErrIDNotAssigned
	}
	err := f.ref.Child(fid.ExportID()).Child("sections").Child(sid.ExportID()).
		Transaction(ctx, func(t db.TransactionNode) (interface{}, error) {
			var v *float64
			if err := t.Unmarshal(&v); err != nil {
				return nil, fmt.Errorf("failed to unmarshal: %w", err)
			}
			if v == nil {
				return index, nil
			} else {
				return nil, exception.ErrAlreadyExists
			}
		})
	if err != nil {
		return fmt.Errorf("failed to add section order: %w", err)
	}
	return nil
}

func (f Form) UpdateSectionOrder(ctx context.Context, fid id.FormID, sid id.SectionID, index float64) error {
	if fid == nil || !fid.HasValue() ||
		sid == nil || !sid.HasValue() {
		return exception.ErrIDNotAssigned
	}
	err := f.ref.Child(fid.ExportID()).Child("sections").Child(sid.ExportID()).
		Transaction(ctx, func(t db.TransactionNode) (interface{}, error) {
			var v *float64
			if err := t.Unmarshal(&v); err != nil {
				return nil, fmt.Errorf("failed to unmarshal: %w", err)
			}
			if v == nil {
				return nil, exception.ErrNotFound
			} else {
				return index, nil
			}
		})
	if err != nil {
		return fmt.Errorf("failed to update section order: %w", err)
	}
	return nil
}
