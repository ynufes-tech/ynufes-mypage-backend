package reader

import (
	"cloud.google.com/go/firestore"
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
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

func (q Question) GetByID(ctx context.Context, id question.ID) (*question.Question, error) {
	var questionEntity entity.Question
	qid := id.ExportID()
	snap, err := q.collection.Doc(qid).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = snap.DataTo(&questionEntity)
	questionEntity.ID = qid
	model, err := questionEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (q Question) GetByEventID(ctx context.Context, id event.ID) ([]question.Question, error) {
	var questions []question.Question
	iter := q.collection.Where("event_id", "==", id.GetValue()).Documents(ctx)
	for {
		var questionEntity entity.Question
		doc, err := iter.Next()
		if err != nil {
			break
		}
		err = doc.DataTo(&questionEntity)
		questionEntity.ID = doc.Ref.ID
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		questions = append(questions, model)
	}
	return questions, nil
}

func (q Question) GetByFormID(ctx context.Context, id event.ID) ([]question.Question, error) {
	var questions []question.Question
	iter := q.collection.Where("form_id", "==", id.GetValue()).Documents(ctx)
	for {
		var questionEntity entity.Question
		doc, err := iter.Next()
		if err != nil {
			break
		}
		err = doc.DataTo(&questionEntity)
		questionEntity.ID = doc.Ref.ID
		model, err := questionEntity.ToModel()
		if err != nil {
			return nil, err
		}
		questions = append(questions, model)
	}
	return questions, nil
}
