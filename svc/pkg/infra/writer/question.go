package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/question"
)

type (
	Question struct {
		ref *db.Ref
	}
)

func NewQuestion(f *firebase.Firebase) Question {
	return Question{
		ref: f.Client(entity.QuestionRootName),
	}
}

func (w Question) Create(ctx context.Context, q *question.Question) error {
	newID := identity.IssueID()
	if q == nil {
		return fmt.Errorf("question is nil")
	}
	if err := (*q).AssignID(newID); err != nil {
		return err
	}
	st, err := (*q).Export()
	if err != nil {
		return fmt.Errorf("failed to export question: %w", err)
	}
	e := entity.NewQuestion(
		(*q).GetID(),
		(*q).GetFormID().ExportID(),
		(*q).GetText(),
		int((*q).GetType()),
		st.Customs,
	)
	if err := w.ref.Child((*q).GetID().ExportID()).
		Set(ctx, e); err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}
	return nil
}

func (w Question) UpdateCustoms(ctx context.Context, id id.QuestionID, customs map[string]interface{}) error {
	if id == nil || !id.HasValue() {
		return exception.ErrIDNotAssigned
	}
	err := w.ref.Child(id.ExportID()).
		Update(ctx, map[string]interface{}{
			"customs": customs,
		})
	if err != nil {
		return fmt.Errorf("failed to update customs: %w", err)
	}
	return nil
}

func (w Question) Set(ctx context.Context, q question.Question) error {
	if q.GetID() == nil || !q.GetID().HasValue() {
		return exception.ErrIDNotAssigned
	}
	st, err := q.Export()
	if err != nil {
		return fmt.Errorf("failed to export question: %w", err)
	}
	e := entity.NewQuestion(
		q.GetID(),
		q.GetFormID().ExportID(),
		q.GetText(),
		int(q.GetType()),
		st.Customs,
	)
	if err := w.ref.Child(q.GetID().ExportID()).
		Set(ctx, e); err != nil {
		return fmt.Errorf("failed to set question: %w", err)
	}
	return nil
}
