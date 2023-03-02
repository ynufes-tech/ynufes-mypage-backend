package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/question"
)

type (
	Question struct {
		collection *firestore.CollectionRef
	}
)

func NewQuestion(client *firestore.Client) Question {
	return Question{
		collection: client.Collection(entity.QuestionCollectionName),
	}
}

func (w Question) Create(ctx context.Context, formID id.FormID, q question.Question) error {
	if !q.GetID().HasValue() {
		return exception.ErrIDNotAssigned
	}
	e := entity.NewQuestion(
		q.GetID(),
		(q.GetEventID()).GetValue(),
		formID.GetValue(),
		q.GetText(),
		int(q.GetType()),
		q.Export().Customs,
	)
	_, err := w.collection.Doc(q.GetID().ExportID()).
		Create(ctx, e)
	if err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}
	return nil
}

func (w Question) UpdateCustoms(ctx context.Context, id id.QuestionID, customs map[string]interface{}) error {
	_, err := w.collection.Doc(id.ExportID()).
		Update(ctx,
			[]firestore.Update{
				{
					Path:  "customs",
					Value: customs,
				},
			})
	if err != nil {
		return fmt.Errorf("failed to update customs: %w", err)
	}
	return nil
}

func (w Question) Set(ctx context.Context, formID id.FormID, q question.Question) error {
	e := entity.NewQuestion(
		q.GetID(),
		(q.GetEventID()).GetValue(),
		formID.GetValue(),
		q.GetText(),
		int(q.GetType()),
		q.Export().Customs,
	)
	_, err := w.collection.Doc(q.GetID().ExportID()).
		Set(ctx, e)
	if err != nil {
		return fmt.Errorf("failed to set question: %w", err)
	}
	return nil
}
