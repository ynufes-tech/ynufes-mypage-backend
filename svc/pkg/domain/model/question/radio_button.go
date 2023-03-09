package question

import (
	"errors"
	"fmt"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	RadioButtonsQuestion struct {
		Basic
		Options      []RadioButtonOption
		OptionsOrder map[RadioButtonOptionID]float64
	}
	RadioButtonOption struct {
		ID   RadioButtonOptionID
		Text string
	}
	RadioButtonOptionID     util.ID
	RadioButtonOptionsOrder map[RadioButtonOptionID]float64
)

const (
	RadioButtonOptionsField      = "options"
	RadioButtonOptionsOrderField = "order"
)

func NewRadioButtonsQuestion(
	id id.QuestionID, text string, options []RadioButtonOption, order RadioButtonOptionsOrder,
	formID id.FormID, sectionID id.SectionID,
) *RadioButtonsQuestion {
	return &RadioButtonsQuestion{
		Basic:        NewBasic(id, text, TypeRadio, formID, sectionID),
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
	optionsOrderData, ok := optionsOrderDataI.(map[string]interface{})
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be []int64 for RadioButtonsQuestion", RadioButtonOptionsOrderField))
	}

	options := make([]RadioButtonOption, 0, len(optionsData))
	optionsOrder := make(map[RadioButtonOptionID]float64, len(optionsOrderData))
	for tid, index := range optionsOrderData {
		fIndex, ok := index.(float64)
		if !ok {
			return nil, errors.New(
				fmt.Sprintf("\"%s\" must be []int64 for RadioButtonsQuestion", RadioButtonOptionsOrderField))
		}
		i, err := identity.ImportID(tid)
		if err != nil {
			return nil, err
		}
		optionsOrder[i] = fIndex
	}

	for oid, text := range optionsData {
		options = append(options, RadioButtonOption{
			ID:   identity.NewID(oid),
			Text: text,
		})
	}
	return NewRadioButtonsQuestion(
		q.ID, q.Text, options, optionsOrder, q.FormID, q.SectionID,
	), nil
}

func (q RadioButtonsQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})

	options := make(map[int64]string, len(q.Options))
	for _, o := range q.Options {
		options[o.ID.GetValue()] = o.Text
	}

	optionsOrder := make(map[string]float64, len(q.OptionsOrder))
	for tid, index := range q.OptionsOrder {
		optionsOrder[tid.ExportID()] = index
	}

	customs[RadioButtonOptionsField] = options
	customs[RadioButtonOptionsOrderField] = optionsOrder

	return NewStandardQuestion(TypeRadio, q.ID, q.FormID, q.SectionID, q.Text, customs)
}

func (q RadioButtonsQuestion) GetOrderedIDs() []RadioButtonOptionID {
	ids := make([]RadioButtonOptionID, 0, len(q.OptionsOrder))
	for oid := range q.OptionsOrder {
		ids = append(ids, oid)
	}
	return ids
}
