package uc

import (
	"context"
	"errors"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type UserInfoUpdateUseCase struct {
	userC command.User
}

type UserInfoUpdateInput struct {
	Ctx     context.Context
	NewUser user.User
}

type UserInfoUpdateOutput struct {
	Error error
}

func NewInfoUpdate(rgst registry.Registry) UserInfoUpdateUseCase {
	return UserInfoUpdateUseCase{
		userC: rgst.Repository().NewUserCommand(),
	}
}

func (uc UserInfoUpdateUseCase) Do(input UserInfoUpdateInput) (*UserInfoUpdateOutput, error) {
	if !input.NewUser.Detail.MeetsBasicRequirement() {
		return &UserInfoUpdateOutput{Error: errors.New("your request does not meet the basic requirement")}, nil
	}
	err := uc.userC.UpdateUserDetail(input.Ctx, input.NewUser.ID, input.NewUser.Detail)
	return &UserInfoUpdateOutput{}, err
}
