package section

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

// SectionFull holds questions themselves
type SectionFull struct {
	Section
	Questions map[id.QuestionID]question.Question
}

func NewSectionFull(
	section Section, questions map[id.QuestionID]question.Question,
) SectionFull {
	return SectionFull{
		Section:   section,
		Questions: questions,
	}
}
