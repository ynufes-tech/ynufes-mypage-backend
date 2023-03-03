package org

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type CreateOrgUseCase struct {
	orgC   command.Org
	eventQ query.Event
}

type CreateOrgInput struct {
	Ctx     context.Context
	EventID id.EventID
	OrgName string
	IsOpen  bool
}

type CreateOrgOutput struct {
	Org org.Org
}

func NewCreateOrg(rgst registry.Registry) CreateOrgUseCase {
	return CreateOrgUseCase{
		orgC:   rgst.Repository().NewOrgCommand(),
		eventQ: rgst.Repository().NewEventQuery(),
	}
}

func (uc CreateOrgUseCase) Do(ipt CreateOrgInput) (*CreateOrgOutput, error) {
	e, err := uc.eventQ.GetByID(ipt.Ctx, ipt.EventID)
	if err != nil {
		return nil, err
	}
	o := org.Org{
		Event:  *e,
		Name:   ipt.OrgName,
		Users:  nil,
		IsOpen: ipt.IsOpen,
	}
	err = uc.orgC.Create(ipt.Ctx, &o)
	if err != nil {
		return nil, err
	}
	return &CreateOrgOutput{
		Org: o,
	}, nil
}
