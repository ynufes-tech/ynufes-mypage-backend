package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/form"
)

type Form struct {
	ref *db.Ref
}

func NewForm(client *firebase.Firebase) *Form {
	return &Form{
		ref: client.Client(entity.FormRootName),
	}
}

func (f Form) GetByID(ctx context.Context, id id.FormID) (*form.Form, error) {
	if !id.HasValue() {
		return nil, exception.ErrIDNotAssigned
	}
	var e entity.Form
	r, err := f.ref.OrderByKey().EqualTo(id.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	if len(r) > 1 {
		fmt.Printf("multiple form found with id: %s", id)
	}
	if err := r[0].Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal form entity: %w", err)
	}
	e.ID = id
	m, err := e.ToModel()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (f Form) ListByEventID(ctx context.Context, eventID id.EventID) ([]form.Form, error) {
	if !eventID.HasValue() {
		return nil, exception.ErrIDNotAssigned
	}
	results, err := f.ref.OrderByChild("event_id").EqualTo(eventID.GetValue()).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}

	forms := make([]form.Form, len(results))
	for i, r := range results {
		var e entity.Form
		if err := r.Unmarshal(&e); err != nil {
			return nil, err
		}
		m, err := e.ToModel()
		if err != nil {
			return nil, fmt.Errorf("failed to convert entity to model in ListByEventID: %w", err)
		}
		forms[i] = *m
	}
	return forms, nil
}
