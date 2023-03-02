package org

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type OrgsUseCase struct {
	orgQ query.Org
	orgC command.Org
}

type OrgsInput struct {
	Ctx    context.Context
	UserID id.UserID
}

type OrgsOutput struct {
	Orgs []org.Org
}

func NewOrgs(rgst registry.Registry) OrgsUseCase {
	return OrgsUseCase{
		orgQ: rgst.Repository().NewOrgQuery(),
		orgC: rgst.Repository().NewOrgCommand(),
	}
}

func (o OrgsUseCase) Do(ipt OrgsInput) (opt *OrgsOutput, err error) {
	orgs, err := o.orgQ.ListByGrantedUserID(ipt.Ctx, ipt.UserID)
	if err != nil {
		return
	}
	return &OrgsOutput{
		Orgs: orgs,
	}, nil
}
