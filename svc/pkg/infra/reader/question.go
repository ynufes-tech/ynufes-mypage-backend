package reader

import (
	"cloud.google.com/go/firestore"
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/question"
)

type Question struct {
	collection *firestore.CollectionRef
}

func NewQuestion(c *firestore.Client) Question {
	return Question{
		collection: c.Collection(entity.QuestionCollectionName),
	}
}

func (q Question) GetByID(ctx context.Context, id id.QuestionID) (*question.Question, error) {
	var questionEntity entity.Question
	qid := id.ExportID()
	snap, err := q.collection.Doc(qid).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = snap.DataTo(&questionEntity)
	questionEntity.ID = id
	model, err := questionEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (q Question) ListByEventID(ctx context.Context, id id.EventID) ([]question.Question, error) {
	var questions []question.Question
	iter := q.collection.Where("event_id", "==", id.GetValue()).Documents(ctx)
	for {
		var questionEntity entity.Question
		doc, err := iter.Next()
		if err != nil {
			break
		}
		err = doc.DataTo(&questionEntity)
		questionEntity.ID = id
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		questions = append(questions, model)
	}
	return questions, nil
}

func (q Question) ListByFormID(ctx context.Context, id id.FormID) ([]question.Question, error) {
	var questions []question.Question
	iter := q.collection.Where("form_id", "==", id.GetValue()).Documents(ctx)
	for {
		var questionEntity entity.Question
		doc, err := iter.Next()
		if err != nil {
			break
		}
		err = doc.DataTo(&questionEntity)
		questionEntity.ID = id
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		questions = append(questions, model)
	}
	return questions, nil
}
