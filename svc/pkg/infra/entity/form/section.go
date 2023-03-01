package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type Section struct {
	ID        int64
	Questions []int64

	// ConditionQuestion a question which determines next section based on its answer
	// Only some of the questions can be condition questions. (e.g. radio, checkbox)
	// If !ConditionQuestion.HasValue(), then proceed to next section
	ConditionQuestion int64

	// ConditionCustoms map[OptionID]NextSectionID
	ConditionCustoms map[string]int64
}

func (s Section) ToModel() (*form.Section, error) {
	qs := make([]question.ID, len(s.Questions))
	for i := range s.Questions {
		qs[i] = identity.NewID(s.Questions[i])
	}

	conditionCustoms := make(map[util.ID]form.SectionID, len(s.ConditionCustoms))
	for k, v := range s.ConditionCustoms {
		i, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		conditionCustoms[i] = identity.NewID(v)
	}

	sec := form.NewSection(
		identity.NewID(s.ID),
		qs,
		identity.NewID(s.ConditionQuestion),
		conditionCustoms,
	)
	return &sec, nil
}
