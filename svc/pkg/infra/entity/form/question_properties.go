package entity

import (
	"errors"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
)

func importQuestionProperties(questionType form.QuestionType, data interface{}) (*form.QuestionProperties, error) {
	var properties form.QuestionProperties
	var err error
	switch questionType {
	case form.CheckBox:
		p, ok := data.(CheckBoxQuestionProperties)
		if !ok {
			return nil, errors.New("invalid data type")
		}
		properties, err = p.Import()
		if err != nil {
			return nil, err
		}
		break
	default:
		return nil, errors.New("invalid question type")
	}
	return &properties, nil
}

type (
	CheckBoxQuestionProperties map[string]struct {
		Text  string `firestore:"text"`
		Order int    `firestore:"order"`
	}
)

func (p CheckBoxQuestionProperties) Import() (form.QuestionProperties, error) {
	var options []form.CheckBoxOption
	for id, data := range p {
		idO, err := identity.ImportID(id)
		if err != nil {
			return nil, err
		}
		options = append(options, form.CheckBoxOption{
			ID:    idO,
			Text:  data.Text,
			Order: data.Order,
		})
	}
	return form.CheckBoxQuestionProperties{
		Options: options,
	}, nil
}

func ExportCheckBoxQuestionProperties(props form.CheckBoxQuestionProperties) (CheckBoxQuestionProperties, error) {
	var options CheckBoxQuestionProperties
	for i := range props.Options {
		if !props.Options[i].ID.HasValue() {
			return nil, errors.New("invalid id")
		}
		d := struct {
			Text  string `firestore:"text"`
			Order int    `firestore:"order"`
		}{
			Text:  props.Options[i].Text,
			Order: props.Options[i].Order,
		}
		options[props.Options[i].ID.ExportID()] = d
	}
	return options, nil
}
