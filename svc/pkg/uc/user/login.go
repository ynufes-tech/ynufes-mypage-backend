package uc

import (
	"context"
	userDomain "ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/domain/service/user"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type LoginUseCase struct {
	userQuery query.User
	login     user.Login
}

type LoginInput struct {
	JWT userDomain.JWT
}

type LoginOutput struct {
	User userDomain.User
}

// NewLoginUseCase TODO: add userQuery, login
func NewLoginUseCase(registry registry.Registry) LoginUseCase {
	return LoginUseCase{}
}

func (uc LoginUseCase) Do(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	userData, err := uc.login.Do(ctx, input.JWT)
	if err != nil {
		return nil, err
	}
	return &LoginOutput{User: userData}, nil
}
