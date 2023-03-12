package section

import (
	"context"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type CreateUseCase struct {
	sectionC command.Section
	formC    command.Form
	formQ    query.Form
}

type CreateInput struct {
	Ctx    context.Context
	UserID id.UserID
	FormID id.FormID
}

type CreateOutput struct {
	SectionID id.SectionID
}

func NewCreate(rgst registry.Registry) CreateUseCase {
	return CreateUseCase{
		sectionC: rgst.Repository().NewSectionCommand(),
		formC:    rgst.Repository().NewFormCommand(),
		formQ:    rgst.Repository().NewFormQuery(),
	}
}

func (uc CreateUseCase) Do(ipt CreateInput) (*CreateOutput, error) {
	targetForm, err := uc.formQ.GetByID(ipt.Ctx, ipt.FormID)
	if err != nil {
		return nil, fmt.Errorf("failed to get form: %w", err)
	}

	targetSection := section.Section{
		FormID:            ipt.FormID,
		QuestionIDs:       nil,
		ConditionQuestion: nil,
		ConditionCustoms:  nil,
	}
	err = uc.sectionC.Create(
		ipt.Ctx,
		&targetSection,
	)
	if err != nil {
		return nil, err
	}
	var newIndex float64
	sections := targetForm.Sections.GetOrderedIDs()
	if len(sections) == 0 {
		newIndex = 1.0
	} else {
		// specify the index as last index
		newIndex = targetForm.Sections[sections[len(sections)-1]] + 1.0
	}
	if err := uc.formC.AddSectionOrder(
		ipt.Ctx, ipt.FormID, targetSection.ID, newIndex,
	); err != nil {
		return nil, fmt.Errorf("failed to add section order: %w", err)
	}
	return &CreateOutput{
		SectionID: targetSection.ID,
	}, nil
}
