package section

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

// Section only holds IDs of questions
type (
	Section struct {
		ID          ID
		QuestionIDs []question.ID

		// ConditionQuestion a question which determines next section based on its answer
		// Only some of the questions can be condition questions. (e.g. radio, checkbox)
		// If !ConditionQuestion.HasValue(), then proceed to next section
		ConditionQuestion question.ID

		// ConditionCustoms map[OptionID]NextSectionID for ConditionQuestion
		ConditionCustoms map[util.ID]ID
	}
	ID util.ID
)

func NewSection(
	id ID,
	questionIDs []question.ID,
	conditionQuestion question.ID,
	conditionCustoms map[util.ID]ID,
) Section {
	return Section{
		ID:                id,
		QuestionIDs:       questionIDs,
		ConditionQuestion: conditionQuestion,
		ConditionCustoms:  conditionCustoms,
	}
}
