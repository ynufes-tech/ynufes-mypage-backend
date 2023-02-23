package question

import (
	"errors"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	ID       util.ID
	Type     int
	Question interface {
		Export() StandardQuestion
		GetType() Type
		GetID() ID
		GetText() string
		GetEventID() event.ID
		GetFormID() form.ID
	}

	StandardQuestion struct {
		ID      ID
		Text    string
		EventID event.ID
		FormID  form.ID
		Type    Type
		Customs map[string]interface{}
	}
)

const (
	TypeCheckBox Type = 1
	TypeRadio    Type = 2
	TypeFile     Type = 3
)

func NewStandardQuestion(t Type, id ID,
	eventID event.ID, formID form.ID, text string, customs map[string]interface{}) StandardQuestion {
	return StandardQuestion{
		ID:      id,
		Text:    text,
		EventID: eventID,
		FormID:  formID,
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
