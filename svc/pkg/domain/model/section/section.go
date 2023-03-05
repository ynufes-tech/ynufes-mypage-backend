package section

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

// Section only holds IDs of questions
type (
	Section struct {
		ID          id.SectionID
		QuestionIDs map[id.QuestionID]float64

		// ConditionQuestion a question which determines next section based on its answer
		// Only some of the questions can be condition questions. (e.g. radio, checkbox)
		// If !ConditionQuestion.HasValue(), then proceed to next section
		ConditionQuestion id.QuestionID

		// ConditionCustoms map[OptionID]NextSectionID for ConditionQuestion
		ConditionCustoms map[util.ID]id.SectionID
	}
)

func NewSection(
	id id.SectionID,
	questionIDs map[id.QuestionID]float64,
	conditionQuestion id.QuestionID,
	conditionCustoms map[util.ID]id.SectionID,
) Section {
	return Section{
		ID:                id,
		QuestionIDs:       questionIDs,
		ConditionQuestion: conditionQuestion,
		ConditionCustoms:  conditionCustoms,
	}
}
