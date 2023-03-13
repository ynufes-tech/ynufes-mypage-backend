package reader

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

type Question struct {
	ref *db.Ref
}

func NewQuestion(f *firebase.Firebase) Question {
	return Question{
		ref: f.Client(entity.QuestionRootName),
	}
}

func (q Question) GetByID(ctx context.Context, qid id.QuestionID) (*question.Question, error) {
	var questionEntity entity.Question
	r, err := q.ref.OrderByKey().
		EqualTo(qid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	if err := r[0].Unmarshal(&questionEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal question entity: %w", err)
	}
	questionEntity.ID = qid
	model, err := questionEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (q Question) ListByEventID(ctx context.Context, eid id.EventID) ([]question.Question, error) {
	results, err := q.ref.OrderByChild("event_id").EqualTo(eid.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	qs := make([]question.Question, len(results))
	for i := range results {
		var questionEntity entity.Question
		if err := results[i].Unmarshal(&questionEntity); err != nil {
			return nil, err
		}
		qid, err := identity.ImportID(results[i].Key())
		if err != nil {
			return nil, fmt.Errorf("failed to import question id from Key(): %w", err)
		}
		questionEntity.ID = qid
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		qs[i] = model
	}
	return qs, nil
}

func (q Question) ListByFormID(ctx context.Context, id id.FormID) ([]question.Question, error) {
	results, err := q.ref.OrderByChild("form_id").EqualTo(id.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	qs := make([]question.Question, len(results))
	for i := range results {
		var questionEntity entity.Question
		if err := results[i].Unmarshal(&questionEntity); err != nil {
			return nil, err
		}
		qid, err := identity.ImportID(results[i].Key())
		if err != nil {
			return nil, fmt.Errorf("failed to import question id from Key(): %w", err)
		}
		questionEntity.ID = qid
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		qs[i] = model
	}
	return qs, nil
}

func (q Question) ListBySectionID(ctx context.Context, id id.SectionID) ([]question.Question, error) {
	results, err := q.ref.OrderByChild("section_id").EqualTo(id.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	qs := make([]question.Question, len(results))
	for i := range results {
		var questionEntity entity.Question
		if err := results[i].Unmarshal(&questionEntity); err != nil {
			return nil, err
		}
		qid, err := identity.ImportID(results[i].Key())
		if err != nil {
			return nil, fmt.Errorf("failed to import question id from Key(): %w", err)
		}
		questionEntity.ID = qid
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		qs[i] = model
	}
	return qs, nil
}
