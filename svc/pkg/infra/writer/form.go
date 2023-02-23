package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/form"
)

type Form struct {
	collection *firestore.CollectionRef
}

func NewForm(client *firestore.Client) *Form {
	return &Form{
		collection: client.Collection(entity.FormCollectionName),
	}
}

func (f Form) Create(ctx context.Context, target form.Form) error {
	eventID := target.EventID.ExportID()
	var roles = make([]int64, len(target.Roles))
	for i := 0; i < len(target.Roles); i++ {
		roles[i] = target.Roles[i].GetValue()
	}
	e := entity.NewForm(
		eventID,
		target.EventID.GetValue(),
		target.Title,
		target.Summary,
		target.Description,
		roles,
		target.Deadline.UnixMilli(),
		target.IsOpen,
	)
	_, err := f.collection.Doc(eventID).Create(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (f Form) Set(ctx context.Context, target form.Form) error {
	eventID := target.EventID.ExportID()
	var roles = make([]int64, len(target.Roles))
	for i := 0; i < len(target.Roles); i++ {
		roles[i] = target.Roles[i].GetValue()
	}
	e := entity.NewForm(
		eventID,
		target.EventID.GetValue(),
		target.Title,
		target.Summary,
		target.Description,
		roles,
		target.Deadline.UnixMilli(),
		target.IsOpen,
	)
	_, err := f.collection.Doc(eventID).Set(ctx, e)
	if err != nil {
		return err
	}
	return nil
}