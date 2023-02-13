package entity

import (
	"errors"
	"fmt"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
)

func ImportQuestion(q Question) (form.Question, error) {
	qid, err := identity.ImportID(q.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to import question id: %w", err)
	}
	basic := form.QuestionBasic{
		QID:           qid,
		QuestionText:  q.QuestionText,
		QuestionOrder: q.Order,
	}
	switch form.QuestionType(q.QuestionType) {
	case form.CheckBox:
		return ImportCheckboxQuestion(q, basic)
	default:
		return nil, fmt.Errorf("unknown question type: %d", q.QuestionType)
	}
}

func ImportCheckboxQuestion(q Question, basic form.QuestionBasic) (*form.CheckboxQuestion, error) {
	var options []form.CheckboxOption
	for k := range q.Properties {
		id, err := identity.ImportID(k)
		if err != nil {
			return nil, fmt.Errorf("failed to import option id: %w", err)
		}
		p, ok := q.Properties[k].(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid option property")
		}
		text, ok := p["text"]
		if !ok {
			return nil, errors.New("invalid option property")
		}
		textStr, ok := text.(string)
		if !ok {
			return nil, errors.New("invalid option property")
		}
		order, ok := p["order"]
		if !ok {
			return nil, errors.New("invalid option property")
		}
		orderInt, ok := order.(int)
		if !ok {
			return nil, errors.New("invalid option property")
		}
		options = append(options, form.CheckboxOption{
			ID:    id,
			Text:  textStr,
			Order: orderInt,
		})
	}
	return &form.CheckboxQuestion{
		QuestionBasic: basic,
		Options:       options,
	}, nil
}
