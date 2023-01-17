package form

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	QuestionProperties interface {
		Type() QuestionType
	}
	CheckBoxQuestionProperties struct {
		Options []CheckBoxOption
	}
	CheckBoxOption struct {
		ID    OptionID
		Text  string
		Order int
	}
	OptionID util.ID
)

func (p CheckBoxQuestionProperties) Type() QuestionType {
	return CheckBox
}

func (q CheckBoxQuestionProperties) Export() interface{} {
	var optStr map[string]interface{}
	for i := range q.Options {
		optStr[q.Options[i].ID.ExportID()] = q.Options[i].Export()
	}
	return optStr
}

func (o CheckBoxOption) Export() interface{} {
	return map[string]interface{}{
		"text":  o.Text,
		"order": o.Order,
	}
}
