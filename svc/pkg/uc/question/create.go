package question

import (
	"context"
	"errors"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type CreateUseCase struct {
	questionC command.Question
	questionQ query.Question
	sectionC  command.Section
	sectionQ  query.Section
}

// CreateInput Order will be ignored if After is not nil or 0 value.
type CreateInput struct {
	Ctx       context.Context
	UserID    id.UserID
	SectionID id.SectionID
	After     id.QuestionID
	Order     *int
	Question  question.Question
}

type CreateOutput struct {
	Question question.Question
}

func NewCreate(rgst registry.Registry) CreateUseCase {
	return CreateUseCase{
		questionC: rgst.Repository().NewQuestionCommand(),
		questionQ: rgst.Repository().NewQuestionQuery(),
		sectionC:  rgst.Repository().NewSectionCommand(),
		sectionQ:  rgst.Repository().NewSectionQuery(),
	}
}

func (u CreateUseCase) Do(ipt CreateInput) (*CreateOutput, error) {
	if err := u.questionC.
		Create(ipt.Ctx, &ipt.Question); err != nil {
		return nil, err
	}

	sec, err := u.sectionQ.GetSectionByID(ipt.Ctx, ipt.SectionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get section in question createUC: %w", err)
	}

	var newIndex float64

	if ipt.After != nil && ipt.After.HasValue() {
		// if after is specified
		newIndex, err = getNewIndexAfter(sec.QuestionIDs, ipt.After)
		if err != nil {
			return nil, err
		}
	} else if ipt.Order != nil {
		// if order is specified
		newIndex, err = getNewIndexPos(sec.QuestionIDs, *ipt.Order)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("after or order must be specified")
	}

	if err := u.sectionC.LinkQuestion(
		ipt.Ctx,
		ipt.SectionID,
		ipt.Question.GetID(),
		newIndex,
	); err != nil {
		return nil, err
	}
	return &CreateOutput{
		Question: ipt.Question,
	}, nil
}

func getNewIndexAfter(qs section.QuestionsOrder, before id.QuestionID) (float64, error) {
	val, has := qs[before]
	if !has {
		return 0, fmt.Errorf("question specified not found in the section")
	}
	qids := qs.GetOrderedIDs()

	for i := range qids {
		if qids[i] != before {
			continue
		}

		// if before is the last question
		if i+1 == len(qids) {
			return val + 1, nil
		}

		return (val + qs[qids[i+1]]) / 2, nil
	}
	return 0, fmt.Errorf("question specified not found in the section")
}

func getNewIndexPos(qs section.QuestionsOrder, pos int) (float64, error) {
	if pos < 0 {
		return 0, fmt.Errorf("invalid position")
	}
	qids := qs.GetOrderedIDs()

	// if first position is specified
	if pos == 0 {
		if len(qids) == 0 {
			return 0, nil
		}
		return qs[qids[0]] - 1, nil
	}

	// if last position is specified
	if len(qids) <= pos {
		return qs[qids[len(qids)-1]] + 1, nil
	}

	return (qs[qids[pos-1]] + qs[qids[pos]]) / 2, nil
}
