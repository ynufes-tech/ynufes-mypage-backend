package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
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

func (q Question) GetByID(ctx context.Context, id id.QuestionID) (*question.Question, error) {
	var questionEntity entity.Question
	if err := q.ref.Child(id.ExportID()).Get(ctx, &questionEntity); err != nil {
		return nil, err
	}
	questionEntity.ID = id
	model, err := questionEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (q Question) ListByEventID(ctx context.Context, id id.EventID) ([]question.Question, error) {
	var questions []question.Question
	results, err := q.ref.OrderByChild("event_id").EqualTo(id.GetValue()).
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
		questionEntity.ID = id
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		qs[i] = model
	}
	return questions, nil
}

func (q Question) ListByFormID(ctx context.Context, id id.FormID) ([]question.Question, error) {
	var questions []question.Question
	results, err := q.ref.OrderByChild("form_id").EqualTo(id.GetValue()).
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
		questionEntity.ID = id
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		qs[i] = model
	}
	return questions, nil
}
