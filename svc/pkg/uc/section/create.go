package section

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type CreateUseCase struct {
	sectionC command.Section
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
	}
}

func (uc CreateUseCase) Do(ipt CreateInput) (*CreateOutput, error) {
	targetSection := section.Section{
		FormID:            ipt.FormID,
		QuestionIDs:       nil,
		ConditionQuestion: nil,
		ConditionCustoms:  nil,
	}
	err := uc.sectionC.Create(
		ipt.Ctx,
		&targetSection,
	)
	if err != nil {
		return nil, err
	}
	return &CreateOutput{
		SectionID: targetSection.ID,
	}, nil
}
