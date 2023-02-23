package question

import (
	"errors"
	"fmt"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	RadioButtonsQuestion struct {
		Basic
		Options      []RadioButtonOption
		OptionsOrder []RadioButtonOptionID
	}
	RadioButtonOption struct {
		ID   RadioButtonOptionID
		Text string
	}
	RadioButtonOptionID util.ID
)

const (
	RadioButtonOptionsField      = "options"
	RadioButtonOptionsOrderField = "order"
)

func NewRadioButtonsQuestion(
	id ID, text string, eventID event.ID, formID form.ID, options []RadioButtonOption, order []RadioButtonOptionID,
) *RadioButtonsQuestion {
	return &RadioButtonsQuestion{
		Basic:        NewBasic(id, text, eventID, formID, TypeRadio),
		Options:      options,
		OptionsOrder: order,
	}
}

func ImportRadioButtonsQuestion(q StandardQuestion) (*RadioButtonsQuestion, error) {
	// check if customs has "options" as map[int64]string, return error if not
	optionsDataI, has := q.Customs[RadioButtonOptionsField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for RadioButtonsQuestion", RadioButtonOptionsField))
	}
	optionsData, ok := optionsDataI.(map[int64]string)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[int64]string for RadioButtonsQuestion", RadioButtonOptionsField))
	}

	// check if customs has "order" as []int64, return error if not
	optionsOrderDataI, has := q.Customs[RadioButtonOptionsOrderField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for RadioButtonsQuestion", RadioButtonOptionsOrderField))
	}
	optionsOrderData, ok := optionsOrderDataI.([]int64)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be []int64 for RadioButtonsQuestion", RadioButtonOptionsOrderField))
	}

	options := make([]RadioButtonOption, 0, len(optionsData))
	optionsOrder := make([]RadioButtonOptionID, 0, len(optionsOrderData))
	for _, id := range optionsOrderData {
		optionsOrder = append(optionsOrder, identity.NewID(id))
	}

	for id, text := range optionsData {
		options = append(options, RadioButtonOption{
			ID:   identity.NewID(id),
			Text: text,
		})
	}
	return NewRadioButtonsQuestion(q.ID, q.Text, q.EventID, q.FormID, options, optionsOrder), nil
}

func (q RadioButtonsQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})

	options := make(map[int64]string, len(q.Options))
	for _, o := range q.Options {
		options[o.ID.GetValue()] = o.Text
	}
	optionsOrder := make([]int64, 0, len(q.OptionsOrder))
	for _, o := range q.OptionsOrder {
		optionsOrder = append(optionsOrder, o.GetValue())
	}

	customs[RadioButtonOptionsField] = options
	customs[RadioButtonOptionsOrderField] = optionsOrder

	return NewStandardQuestion(TypeRadio, q.ID, q.EventID, q.FormID, q.Text, customs)
}
