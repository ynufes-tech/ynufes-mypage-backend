package question

import (
	"errors"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	Type     int
	Question interface {
		Export() (*StandardQuestion, error)
		GetType() Type
		AssignID(id.QuestionID) error
		GetID() id.QuestionID
		GetText() string
		GetFormID() id.FormID
	}

	StandardQuestion struct {
		ID      id.QuestionID
		Text    string
		FormID  id.FormID
		Type    Type
		Customs map[string]interface{}
	}
)

const (
	TypeCheckBox Type = 1
	TypeRadio    Type = 2
	TypeFile     Type = 3
)

func (t Type) String() string {
	switch t {
	case TypeCheckBox:
		return "checkbox"
	case TypeRadio:
		return "radio"
	case TypeFile:
		return "file"
	default:
		return "unknown"
	}
}

func NewType(t string) (Type, error) {
	switch t {
	case "checkbox":
		return TypeCheckBox, nil
	case "radio":
		return TypeRadio, nil
	case "file":
		return TypeFile, nil
	default:
		return 0, errors.New("invalid question type")
	}
}

func NewStandardQuestion(
	t Type, id id.QuestionID, formID id.FormID, text string, customs map[string]interface{},
) *StandardQuestion {
	return &StandardQuestion{
		ID:      id,
		Text:    text,
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
