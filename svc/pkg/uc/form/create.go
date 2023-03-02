package form

import (
	"context"
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type CreateUseCase struct {
	eventQ query.Event
	formC  command.Form
}

type CreateInput struct {
	Ctx         context.Context
	User        user.User
	EventID     id.EventID
	Title       string
	Summary     string
	Description string
	Deadline    time.Time
}

type CreateOutput struct {
	FormID id.FormID
}

func NewCreate(rgst registry.Registry) *CreateUseCase {
	return &CreateUseCase{
		eventQ: rgst.Repository().NewEventQuery(),
		formC:  rgst.Repository().NewFormCommand(),
	}
}

func (uc CreateUseCase) Do(ipt CreateInput) (*CreateOutput, error) {
	_, err := uc.eventQ.GetByID(ipt.Ctx, ipt.EventID)
	if err != nil {
		return nil, err
	}
	// TODO: check if the agent has appropriate role.

	formID := identity.IssueID()
	if err := uc.formC.Create(ipt.Ctx,
		form.Form{
			ID:          formID,
			EventID:     ipt.EventID,
			Title:       ipt.Title,
			Summary:     ipt.Summary,
			Description: ipt.Description,
			// TODO: add roles
			Roles:    nil,
			Deadline: ipt.Deadline,
			IsOpen:   false,
		}); err != nil {
		return nil, err
	}
	return &CreateOutput{
		FormID: formID,
	}, nil
}
