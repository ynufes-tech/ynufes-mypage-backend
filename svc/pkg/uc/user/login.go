package uc

import (
	"context"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/jwt"
	"ynufes-mypage-backend/pkg/setting"
	userDomain "ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type LoginUseCase struct {
	userQuery query.User
	jwtSecret string
}

type LoginInput struct {
	JWT userDomain.JWT
}

type LoginOutput struct {
	User userDomain.User
}

func NewLoginUseCase(registry registry.Registry) LoginUseCase {
	config := setting.Get()
	return LoginUseCase{
		userQuery: registry.Repository().NewUserQuery(),
		jwtSecret: config.Application.Admin.JwtSecret,
	}
}

func (uc LoginUseCase) Do(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	claims, err := jwt.Verify(string(input.JWT), uc.jwtSecret)
	if err != nil {
		return nil, err
	}
	id, err := identity.ImportID(claims.Id)
	if err != nil {
		return nil, err
	}
	userData, err := uc.userQuery.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &LoginOutput{
		User: *userData,
	}, nil
}
