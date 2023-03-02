package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type Section struct {
	ID        int64   `json:"id"`
	Questions []int64 `json:"questions"`

	// ConditionQuestion a question which determines next section based on its answer
	// Only some of the questions can be condition questions. (e.g. radio, checkbox)
	// If !ConditionQuestion.HasValue(), then proceed to next section
	ConditionQuestion int64 `json:"c_question"`

	// ConditionCustoms map[OptionID]NextSectionID
	ConditionCustoms map[string]int64 `json:"c_customs"`
}

func NewSection(
	id int64,
	qs []int64,
	cq int64,
	cc map[string]int64,
) Section {
	return Section{
		ID:                id,
		Questions:         qs,
		ConditionQuestion: cq,
		ConditionCustoms:  cc,
	}
}

func (s Section) ToModel() (*section.Section, error) {
	qs := make([]id.QuestionID, len(s.Questions))
	for i := range s.Questions {
		qs[i] = identity.NewID(s.Questions[i])
	}

	conditionCustoms := make(map[util.ID]id.SectionID, len(s.ConditionCustoms))
	for k, v := range s.ConditionCustoms {
		i, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		conditionCustoms[i] = identity.NewID(v)
	}

	sec := section.NewSection(
		identity.NewID(s.ID),
		qs,
		identity.NewID(s.ConditionQuestion),
		conditionCustoms,
	)
	return &sec, nil
}
