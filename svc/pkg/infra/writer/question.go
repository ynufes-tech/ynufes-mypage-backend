package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/question"
)

type (
	Question struct {
		collection *firestore.CollectionRef
	}
)

func NewQuestion(client *firestore.Client) *Question {
	return &Question{
		collection: client.Collection(entity.QuestionCollectionName),
	}
}

func (w Question) Create(ctx context.Context, q question.Question) error {
	e := entity.Question{
		EventID: (q.GetEventID()).GetValue(),
		FormID:  (q.GetFormID()).GetValue(),
		Text:    q.GetText(),
		Type:    int(q.GetType()),
		Customs: q.Export().Customs,
	}
	_, err := w.collection.Doc(q.GetID().ExportID()).
		Create(ctx, e)
	if err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}
	return nil
}

func (w Question) UpdateCustoms(ctx context.Context, q question.Question) error {
	_, err := w.collection.Doc(q.GetID().ExportID()).
		Update(ctx,
			[]firestore.Update{
				{
					Path:  "customs",
					Value: q.Export().Customs,
				},
			})
	if err != nil {
		return fmt.Errorf("failed to update customs: %w", err)
	}
	return nil
}

func (w Question) Set(ctx context.Context, q question.Question) error {
	e := entity.Question{
		EventID: (q.GetEventID()).GetValue(),
		FormID:  (q.GetFormID()).GetValue(),
		Text:    q.GetText(),
		Type:    int(q.GetType()),
		Customs: q.Export().Customs,
	}
	_, err := w.collection.Doc(q.GetID().ExportID()).
		Set(ctx, e)
	if err != nil {
		return fmt.Errorf("failed to set question: %w", err)
	}
	return nil
}
