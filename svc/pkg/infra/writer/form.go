package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
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
	if target.ID.HasValue() {
		return exception.ErrIDAlreadyAssigned
	}
	tid := identity.IssueID()
	sections := make(map[string]bool, len(target.Sections))
	for i := range target.Sections {
		sections[target.Sections[i].ExportID()] = true
	}

	var roles = make(map[string]bool, len(target.Roles))
	for i := 0; i < len(target.Roles); i++ {
		roles[target.Roles[i].ExportID()] = true
	}
	e := entity.NewForm(
		tid,
		target.EventID.GetValue(),
		target.Title,
		target.Summary,
		target.Description,
		sections,
		roles,
		target.Deadline.UnixMilli(),
		target.IsOpen,
	)
	err := f.ref.Child(target.ID.ExportID()).Set(ctx, e)
	if err != nil {
		return err
	}
	target.ID = tid
	return nil
}

func (f Form) Set(ctx context.Context, target form.Form) error {
	if !target.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	sections := make(map[string]bool, len(target.Sections))
	for i := range target.Sections {
		sections[target.Sections[i].ExportID()] = true
	}

	var roles = make(map[string]bool, len(target.Roles))
	for i := 0; i < len(target.Roles); i++ {
		roles[target.Roles[i].ExportID()] = true
	}
	e := entity.NewForm(
		target.ID,
		target.EventID.GetValue(),
		target.Title,
		target.Summary,
		target.Description,
		sections,
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
