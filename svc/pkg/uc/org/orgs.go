package org

import (
	"context"
	"errors"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type OrgsUseCase struct {
	orgQ      query.Org
	orgC      command.Org
	relationQ query.Relation
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
		orgQ:      rgst.Repository().NewOrgQuery(),
		orgC:      rgst.Repository().NewOrgCommand(),
		relationQ: rgst.Repository().NewRelationQuery(),
	}
}

func (o OrgsUseCase) Do(ipt OrgsInput) (opt *OrgsOutput, err error) {
	orgIDs, err := o.relationQ.ListOrgIDsByUserID(ipt.Ctx, ipt.UserID)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return &OrgsOutput{
				Orgs: []org.Org{},
			}, nil
		}
		return nil, err
	}
	orgs := make([]org.Org, 0, len(orgIDs))
	for i := 0; i < len(orgIDs); i++ {
		t, err := o.orgQ.GetByID(ipt.Ctx, orgIDs[i])
		if err != nil {
			if errors.Is(err, exception.ErrNotFound) {
				fmt.Println("org should be found, but not found")
				continue
			}
			return nil, err
		}
		orgs = append(orgs, *t)
	}
	return &OrgsOutput{
		Orgs: orgs,
	}, nil
}
