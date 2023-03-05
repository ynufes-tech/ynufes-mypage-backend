package entity

import (
	"sort"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

const SectionRootName = "Sections"

type Section struct {
	ID string `json:"-"`

	// Questions map[QID]order
	// One of the idea to manage order is to apply fractional indexing,
	// although implementing here may make data structure more complicated,
	// decided not to implement.
	Questions map[string]int `json:"questions"`

	// ConditionQuestion a question which determines next section based on its answer
	// Only some of the questions can be condition questions. (e.g. radio, checkbox)
	// If !ConditionQuestion.HasValue(), then proceed to next section
	ConditionQuestion string `json:"c_question"`

	// ConditionCustoms map[OptionID]NextSectionID
	ConditionCustoms map[string]string `json:"c_customs"`
}

func NewSection(
	sectionID string,
	questions map[string]int,
	conditionQID string,
	conditionCustoms map[string]string,
) Section {
	return Section{
		ID:                sectionID,
		Questions:         questions,
		ConditionQuestion: conditionQID,
		ConditionCustoms:  conditionCustoms,
	}
}

func (s Section) ToModel() (*section.Section, error) {
	qs, err := sortQuestions(s.Questions)
	if err != nil {
		return nil, err
	}

	conditionCustoms := make(map[util.ID]id.SectionID, len(s.ConditionCustoms))
	for k, v := range s.ConditionCustoms {
		i, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		nextS, err := identity.ImportID(v)
		if err != nil {
			return nil, err
		}
		conditionCustoms[i] = nextS
	}

	sid, err := identity.ImportID(s.ID)
	if err != nil {
		return nil, err
	}

	cid, err := identity.ImportID(s.ConditionQuestion)
	if err != nil {
		return nil, err
	}

	sec := section.NewSection(
		sid,
		qs,
		cid,
		conditionCustoms,
	)
	return &sec, nil
}

type question struct {
	order int
	id    id.QuestionID
}

type questions []question

func sortQuestions(qs map[string]int) ([]id.QuestionID, error) {
	q, err := newQuestions(qs)
	if err != nil {
		return nil, err
	}
	sort.Sort(q)
	ids := make([]id.QuestionID, len(q))
	for i := range q {
		ids[i] = q[i].id
	}
	return ids, nil
}

func newQuestions(qs map[string]int) (questions, error) {
	q := make(questions, 0, len(qs))
	for k, v := range qs {
		tid, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		q = append(q, question{
			order: v,
			id:    tid,
		})
	}
	return q, nil
}

func (q questions) Len() int {
	return len(q)
}

func (q questions) Less(i, j int) bool {
	return q[i].order < q[j].order
}

func (q questions) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
