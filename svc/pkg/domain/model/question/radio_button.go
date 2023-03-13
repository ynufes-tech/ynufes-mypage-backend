package question

import (
	"errors"
	"fmt"
	"sort"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/typecast"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	RadioButtonsQuestion struct {
		Basic
		Options      map[RadioButtonOptionID]RadioButtonOption
		OptionsOrder RadioButtonOptionsOrder
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
	id id.QuestionID, text string, options map[RadioButtonOptionID]RadioButtonOption, order RadioButtonOptionsOrder, formID id.FormID,
) *RadioButtonsQuestion {
	return &RadioButtonsQuestion{
		Basic:        NewBasic(id, text, TypeRadio, formID),
		Options:      options,
		OptionsOrder: order,
	}
}

func ImportRadioButtonsQuestion(q StandardQuestion) (*RadioButtonsQuestion, error) {
	optionsDataI, has := q.Customs[RadioButtonOptionsField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for RadioButtonsQuestion", RadioButtonOptionsField))
	}
	optionsData, err := typecast.ConvertToStringMapString(optionsDataI)
	if err != nil {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]string for RadioButtonsQuestion", RadioButtonOptionsField))
	}

	optionsOrderDataI, has := q.Customs[RadioButtonOptionsOrderField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for RadioButtonsQuestion", RadioButtonOptionsOrderField))
	}
	optionsOrderData, err := typecast.ConvertToStringMapFloat64(optionsOrderDataI)
	if err != nil {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]float64 for RadioButtonsQuestion", RadioButtonOptionsOrderField))
	}

	options := make(map[RadioButtonOptionID]RadioButtonOption, len(optionsData))
	optionsOrder := make(map[RadioButtonOptionID]float64, len(optionsOrderData))
	for tid, index := range optionsOrderData {
		i, err := identity.ImportID(tid)
		if err != nil {
			return nil, err
		}
		optionsOrder[i] = index
	}

	for oid, text := range optionsData {
		i, err := identity.ImportID(oid)
		if err != nil {
			return nil, err
		}
		options[i] = RadioButtonOption{
			ID:   i,
			Text: text,
		}
	}
	return NewRadioButtonsQuestion(
		q.ID, q.Text, options, optionsOrder, q.FormID,
	), nil
}

func (q RadioButtonsQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})

	options := make(map[string]string, len(q.Options))
	for _, o := range q.Options {
		options[o.ID.ExportID()] = o.Text
	}

	optionsOrder := make(map[string]float64, len(q.OptionsOrder))
	for tid, index := range q.OptionsOrder {
		optionsOrder[tid.ExportID()] = index
	}

	customs[RadioButtonOptionsField] = options
	customs[RadioButtonOptionsOrderField] = optionsOrder

	return NewStandardQuestion(TypeRadio, q.ID, q.FormID, q.Text, customs)
}

func (o RadioButtonOptionsOrder) GetOrderedIDs() []RadioButtonOptionID {
	ids := make([]RadioButtonOptionID, 0, len(o))
	for oid := range o {
		ids = append(ids, oid)
	}
	sort.Slice(ids, func(i, j int) bool {
		return o[ids[i]] < o[ids[j]]
	})
	return ids
}
