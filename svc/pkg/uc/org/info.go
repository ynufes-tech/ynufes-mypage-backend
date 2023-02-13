package org

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type InfoUseCase struct {
	orgQ query.Org
}

type InfoInput struct {
	Ctx context.Context
	ID  org.ID
}

type InfoOutput struct {
	Org org.Org
}

func NewInfo(rgst registry.Registry) InfoUseCase {
	return InfoUseCase{
		orgQ: rgst.Repository().NewOrgQuery(),
	}
}

func (uc InfoUseCase) Do(ipt InfoInput) (*InfoOutput, error) {
	o, err := uc.orgQ.GetByID(ipt.Ctx, ipt.ID)
	if err != nil {
		return nil, err
	}
	return &InfoOutput{
		Org: *o,
	}, nil
}
