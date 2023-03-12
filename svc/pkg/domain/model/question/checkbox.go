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
	CheckBoxOptionID util.ID

	CheckBoxQuestion struct {
		Basic
		Options      map[CheckBoxOptionID]CheckBoxOption
		OptionsOrder CheckboxOptionsOrder
	}

	CheckBoxOption struct {
		ID   CheckBoxOptionID
		Text string
	}

	CheckboxOptionsOrder map[CheckBoxOptionID]float64
)

const (
	CheckBoxOptionsField      = "options"
	CheckBoxOptionsOrderField = "order"
)

func NewCheckBoxQuestion(
	id id.QuestionID, text string, options map[CheckBoxOptionID]CheckBoxOption, optionsOrder CheckboxOptionsOrder, formID id.FormID,
) *CheckBoxQuestion {
	return &CheckBoxQuestion{
		Basic:        NewBasic(id, text, TypeCheckBox, formID),
		Options:      options,
		OptionsOrder: optionsOrder,
	}
}

func ImportCheckBoxQuestion(q StandardQuestion) (*CheckBoxQuestion, error) {
	// CheckBoxOptionsField should be map[string]string

	optionsDataI, has := q.Customs[CheckBoxOptionsField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for CheckBoxQuestion", CheckBoxOptionsField))
	}
	optionsData, err := typecast.ConvertToStringMapString(optionsDataI)
	if err != nil {
		fmt.Printf("optionsDataI: %#v %T\n", optionsDataI, optionsDataI)
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]string for CheckBoxQuestion", CheckBoxOptionsField))
	}

	// check if customs has "order" as map[string]float, return error if not
	optionsOrderDataI, has := q.Customs[CheckBoxOptionsOrderField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for CheckBoxQuestion", CheckBoxOptionsOrderField))
	}
	optionsOrderData, err := typecast.ConvertToStringMapFloat64(optionsOrderDataI)
	if err != nil {
		fmt.Printf("optionsOrderDataI: %#v %T\n", optionsOrderDataI, optionsOrderDataI)
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]float64 for CheckBoxQuestion", CheckBoxOptionsOrderField))
	}

	options := make(map[CheckBoxOptionID]CheckBoxOption, len(optionsData))
	optionsOrder := make(map[CheckBoxOptionID]float64, len(optionsOrderData))
	for oid, index := range optionsOrderData {
		oid, err := identity.ImportID(oid)
		if err != nil {
			return nil, err
		}
		optionsOrder[oid] = index
	}

	for oid, text := range optionsData {
		i, err := identity.ImportID(oid)
		if err != nil {
			return nil, err
		}
		options[i] = CheckBoxOption{
			ID:   i,
			Text: text,
		}
	}
	return NewCheckBoxQuestion(q.ID, q.Text, options, optionsOrder, q.FormID), nil
}

func (q CheckBoxQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})
	options := make(map[string]string, len(q.Options))
	for _, option := range q.Options {
		options[option.ID.ExportID()] = option.Text
	}
	optionsOrder := make(map[string]float64, len(q.OptionsOrder))
	for oid, index := range q.OptionsOrder {
		optionsOrder[oid.ExportID()] = index
	}
	customs[CheckBoxOptionsField] = options
	customs[CheckBoxOptionsOrderField] = optionsOrder
	return NewStandardQuestion(TypeCheckBox, q.ID, q.FormID, q.Text, customs)
}

func (o CheckboxOptionsOrder) GetOrderedIDs() []CheckBoxOptionID {
	ids := make([]CheckBoxOptionID, 0, len(o))
	for oid := range o {
		ids = append(ids, oid)
	}
	sort.Slice(ids, func(i, j int) bool {
		return o[ids[i]] < o[ids[j]]
	})
	return ids
}
