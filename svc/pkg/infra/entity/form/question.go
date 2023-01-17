package entity

import (
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

func ImportCheckboxQuestion(q Question, basic form.QuestionBasic) (form.CheckboxQuestion, error) {
	// TODO: implement Options converter
	return form.CheckboxQuestion{
		QuestionBasic: basic,
		Options:       nil,
	}, nil
}
