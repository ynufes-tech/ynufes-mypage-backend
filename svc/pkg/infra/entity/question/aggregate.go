package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

const QuestionRootName = "Questions"

type Question struct {
	ID      id.QuestionID          `json:"-"`
	EventID int64                  `json:"event_id"`
	FormID  int64                  `json:"form_id"`
	Text    string                 `json:"text"`
	Type    int                    `json:"type"`
	Customs map[string]interface{} `json:"customs"`
}

func NewQuestion(
	id id.QuestionID,
	eventID, formID int64,
	text string,
	qType int,
	customs map[string]interface{},
) Question {
	return Question{
		ID:      id,
		EventID: eventID,
		FormID:  formID,
		Text:    text,
		Type:    qType,
		Customs: customs,
	}
}

func (q Question) ToModel() (question.Question, error) {
	sq := question.NewStandardQuestion(
		question.Type(q.Type),
		q.ID,
		identity.NewID(q.EventID),
		q.Text,
		q.Customs,
	)
	return sq.ToQuestion()
}
