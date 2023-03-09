package section

import (
	"sort"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

// Section only holds IDs of questions
type (
	Section struct {
		ID          id.SectionID
		FormID      id.FormID
		QuestionIDs QuestionsOrder

		// ConditionQuestion a question which determines next section based on its answer
		// Only some of the questions can be condition questions. (e.g. radio, checkbox)
		// If !ConditionQuestion.HasValue(), then proceed to next section
		ConditionQuestion id.QuestionID

		// ConditionCustoms map[OptionID]NextSectionID for ConditionQuestion
		ConditionCustoms map[util.ID]id.SectionID
	}

	// QuestionsOrder map[QuestionID]orderValue
	// order of questions are managed by fractional indexing.
	QuestionsOrder map[id.QuestionID]float64
)

func NewSection(
	id id.SectionID,
	formID id.FormID,
	questionIDs map[id.QuestionID]float64,
	conditionQuestion id.QuestionID,
	conditionCustoms map[util.ID]id.SectionID,
) Section {
	return Section{
		ID:                id,
		FormID:            formID,
		QuestionIDs:       questionIDs,
		ConditionQuestion: conditionQuestion,
		ConditionCustoms:  conditionCustoms,
	}
}

type qEntry struct {
	qid id.QuestionID
	idx float64
}
type qEntries []qEntry

func (o QuestionsOrder) GetOrderedIDs() []id.QuestionID {
	ids := make([]qEntry, 0, len(o))
	for tid := range o {
		ids = append(ids, qEntry{tid, o[tid]})
	}
	sort.Sort(qEntries(ids))
	ordered := make([]id.QuestionID, 0, len(o))
	for _, tid := range ids {
		ordered = append(ordered, tid.qid)
	}
	return ordered
}

func (q qEntries) Len() int {
	return len(q)
}

func (q qEntries) Less(i, j int) bool {
	return q[i].idx < q[j].idx
}

func (q qEntries) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
