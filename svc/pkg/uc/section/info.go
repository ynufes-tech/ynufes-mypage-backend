package section

import (
	"context"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type InfoUseCase struct {
	sectionQ  query.Section
	questionQ query.Question
}

type InfoInput struct {
	Ctx       context.Context
	UserID    id.UserID
	SectionID id.SectionID
}

type InfoOutput struct {
	Section section.SectionFull
}

func NewInfo(rgst registry.Registry) InfoUseCase {
	return InfoUseCase{
		sectionQ:  rgst.Repository().NewSectionQuery(),
		questionQ: rgst.Repository().NewQuestionQuery(),
	}
}

func (uc InfoUseCase) Do(ipt InfoInput) (*InfoOutput, error) {
	s, err := uc.sectionQ.GetSectionByID(ipt.Ctx, ipt.SectionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get section: %w", err)
	}

	questions, err := uc.questionQ.ListByFormID(ipt.Ctx, s.FormID)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions: %w", err)
	}

	qMap := make(map[id.QuestionID]question.Question, len(questions))
	for i := range questions {
		qMap[questions[i].GetID()] = questions[i]
	}

	return &InfoOutput{
		Section: section.NewSectionFull(*s, qMap),
	}, nil
}
