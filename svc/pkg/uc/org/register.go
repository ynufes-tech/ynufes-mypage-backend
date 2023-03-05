package org

import (
	"context"
	"errors"
	"fmt"
	"log"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/jwt"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type RegisterUseCase struct {
	orgC      command.Org
	orgQ      query.Org
	relationC command.Relation
	relationQ query.Relation
	jwtSecret string
}

type RegisterInput struct {
	Ctx    context.Context
	UserID id.UserID
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
		relationC: rgst.Repository().NewRelationCommand(),
		relationQ: rgst.Repository().NewRelationQuery(),
		jwtSecret: config.Application.Authentication.JwtSecret,
	}
}

func (uc RegisterUseCase) Do(ipt RegisterInput) (*RegisterOutput, error) {
	verify, err := jwt.Verify(ipt.Token, uc.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token in RegisterUC: %w", err)
	}
	orgID, err := identity.ImportID(verify.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to import orgID in RegisterUC: %w", err)
	}
	o, err := uc.orgQ.GetByID(ipt.Ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("failed to get org in RegisterUC: %w", err)
	}

	if uc.hasUser(ipt.Ctx, orgID, ipt.UserID) {
		return &RegisterOutput{
			Added: false,
			Org:   *o,
		}, nil
	}
	if err := uc.relationC.
		CreateOrgUser(ipt.Ctx, orgID, ipt.UserID); err != nil {
		return nil, err
	}
	return &RegisterOutput{
		Added: true,
		Org:   *o,
	}, nil
}

func (uc RegisterUseCase) hasUser(ctx context.Context, targetOrg id.OrgID, targetUser id.UserID) bool {
	orgIDs, err := uc.relationQ.ListOrgIDsByUserID(ctx, targetUser)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return false
		}
		log.Println(err)
		return false
	}
	return orgIDs.HasOrgID(targetOrg)
}
