package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

const QuestionCollectionName = "Questions"

type Question struct {
	ID      string                 `firestore:"-"`
	EventID int64                  `firestore:"event_id"`
	FormID  int64                  `firestore:"form_id"`
	Text    string                 `firestore:"text"`
	Type    int                    `firestore:"type"`
	Customs map[string]interface{} `firestore:"customs"`
}

func (q Question) ToModel() (question.Question, error) {
	id, err := identity.ImportID(q.ID)
	if err != nil {
		return nil, err
	}
	sq := question.NewStandardQuestion(
		question.Type(q.Type),
		id,
		identity.NewID(q.EventID),
		identity.NewID(q.FormID),
		q.Text,
		q.Customs,
	)
	return sq.ToQuestion()
}
