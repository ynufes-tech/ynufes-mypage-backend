package question

import (
	"errors"
	"fmt"
	"sort"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	CheckBoxOptionID util.ID

	CheckBoxQuestion struct {
		Basic
		Options      []CheckBoxOption
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
	id id.QuestionID, text string, options []CheckBoxOption, optionsOrder CheckboxOptionsOrder,
	formID id.FormID, sectionID id.SectionID,
) *CheckBoxQuestion {
	return &CheckBoxQuestion{
		Basic:        NewBasic(id, text, TypeCheckBox, formID, sectionID),
		Options:      options,
		OptionsOrder: optionsOrder,
	}
}

func ImportCheckBoxQuestion(q StandardQuestion) (*CheckBoxQuestion, error) {
	// CheckBoxOptionsField should be map[string]string

	// Although you cannot cast map[string]interface{} to map[string]string,
	// you have to iterate over the map and cast each value to string.
	// First, check if customs has CheckBoxOptionsField as map[string]interface{}, return error if not.
	optionsDataI, has := q.Customs[CheckBoxOptionsField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for CheckBoxQuestion", CheckBoxOptionsField))
	}
	optionsData, ok := optionsDataI.(map[string]string)
	if !ok {
		fmt.Printf("optionsDataI: %T\n", optionsDataI)
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]string for CheckBoxQuestion", CheckBoxOptionsField))
	}

	// check if customs has "order" as []int64, return error if not
	optionsOrderDataI, has := q.Customs[CheckBoxOptionsOrderField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for CheckBoxQuestion", CheckBoxOptionsOrderField))
	}
	optionsOrderData, ok := optionsOrderDataI.(map[string]float64)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be []int64 for CheckBoxQuestion", CheckBoxOptionsOrderField))
	}

	options := make([]CheckBoxOption, 0, len(optionsData))
	optionsOrder := make(map[CheckBoxOptionID]float64, len(optionsOrderData))
	for oid, index := range optionsOrderData {
		if !ok {
			return nil, errors.New(
				fmt.Sprintf("Option order must be int64 for CheckBoxQuestion"))
		}
		oid, err := identity.ImportID(oid)
		if err != nil {
			return nil, err
		}
		optionsOrder[oid] = index
	}

	for oid, text := range optionsData {
		// here we cast textI to string
		if !ok {
			return nil, errors.New(
				fmt.Sprintf("Option text must be string for CheckBoxQuestion"))
		}
		i, err := identity.ImportID(oid)
		if err != nil {
			return nil, err
		}
		options = append(options, CheckBoxOption{
			ID:   i,
			Text: text,
		})
	}
	return NewCheckBoxQuestion(q.ID, q.Text, options, optionsOrder, q.FormID, q.SectionID), nil
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
	return NewStandardQuestion(TypeCheckBox, q.ID, q.FormID, q.SectionID, q.Text, customs)
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
