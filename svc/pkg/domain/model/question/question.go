package question

import (
	"errors"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	Type     int
	Question interface {
		Export() StandardQuestion
		GetType() Type
		AssignID(id.QuestionID) error
		GetID() id.QuestionID
		GetText() string
		GetEventID() id.EventID
		GetFormID() id.FormID
	}

	StandardQuestion struct {
		ID      id.QuestionID
		Text    string
		EventID id.EventID
		Type    Type
		Customs map[string]interface{}
	}
)

const (
	TypeCheckBox Type = 1
	TypeRadio    Type = 2
	TypeFile     Type = 3
)

func NewStandardQuestion(t Type, id id.QuestionID,
	eventID id.EventID, text string, customs map[string]interface{}) StandardQuestion {
	return StandardQuestion{
		ID:      id,
		Text:    text,
		EventID: eventID,
		Type:    t,
		Customs: customs,
	}
}

func (q StandardQuestion) ToQuestion() (Question, error) {
	switch q.Type {
	case TypeCheckBox:
		return ImportCheckBoxQuestion(q)
	case TypeRadio:
		return ImportRadioButtonsQuestion(q)
	case TypeFile:
		return ImportFileQuestion(q)
	default:
		return nil, errors.New("invalid question type")
	}
}
