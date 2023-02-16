package org

import (
	"context"
	"fmt"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/jwt"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type RegisterUseCase struct {
	orgC      command.Org
	orgQ      query.Org
	jwtSecret string
}

type RegisterInput struct {
	Ctx    context.Context
	UserID user.ID
	Token  string
}

type RegisterOutput struct {
	Added bool
	Org   org.Org
}

func NewRegister(rgst registry.Registry) RegisterUseCase {
	config := setting.Get()
	return RegisterUseCase{
		orgC:      rgst.Repository().NewOrgCommand(),
		orgQ:      rgst.Repository().NewOrgQuery(),
		jwtSecret: config.Application.Authentication.JwtSecret,
	}
}

func (uc RegisterUseCase) Do(ipt RegisterInput) (*RegisterOutput, error) {
	verify, err := jwt.Verify(ipt.Token, uc.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token in RegisterUC: %w", err)
	}
	id, err := identity.ImportID(verify.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to import id in RegisterUC: %w", err)
	}
	o, err := uc.orgQ.GetByID(ipt.Ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get org in RegisterUC: %w", err)
	}
	if hasUser(&o.Users, ipt.UserID) {
		return &RegisterOutput{
			Added: false,
			Org:   *o,
		}, nil
	}
	o.Users = append(o.Users, ipt.UserID)
	if err := uc.orgC.UpdateUsers(ipt.Ctx, *o); err != nil {
		return nil, err
	}
	return &RegisterOutput{
		Added: true,
		Org:   *o,
	}, nil
}

func hasUser(users *[]user.ID, targetUser user.ID) bool {
	for _, m := range *users {
		if m.GetValue() == targetUser.GetValue() {
			return true
		}
	}
	return false
}
