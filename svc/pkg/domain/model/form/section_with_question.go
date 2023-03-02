package form

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

// SectionFull holds questions themselves
type SectionFull struct {
	Section
	Questions map[question.ID]question.Question
}

func NewSectionFull(
	section Section, questions map[question.ID]question.Question,
) SectionFull {
	return SectionFull{
		Section:   section,
		Questions: questions,
	}
}
