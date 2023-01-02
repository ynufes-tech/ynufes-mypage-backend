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
	Ctx       context.Context
	OldUser   *user.User
	NewDetail user.Detail
}

type UserInfoUpdateOutput struct {
	Error error
}

func NewInfoUpdate(rgst registry.Registry) UserInfoUpdateUseCase {
	return UserInfoUpdateUseCase{
		userC: rgst.Repository().NewUserCommand(),
	}
}

func (uc UserInfoUpdateUseCase) Do(input UserInfoUpdateInput) UserInfoUpdateOutput {
	if !input.NewDetail.MeetsBasicRequirement() {
		return UserInfoUpdateOutput{Error: errors.New("invalid user update request")}
	}
	err := uc.userC.UpdateUserDetail(input.Ctx, input.OldUser, input.NewDetail)
	return UserInfoUpdateOutput{
		Error: err,
	}
}
