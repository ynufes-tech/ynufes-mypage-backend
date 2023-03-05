package form

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/domain/service/access"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type InfoUseCase struct {
	formQ  query.Form
	access access.AccessController
}

type InfoInput struct {
	Ctx    context.Context
	UserID id.UserID
	OrgID  id.OrgID
	FormID id.FormID
}

type InfoOutput struct {
	Form form.Form
}

func NewInfo(rgst registry.Registry) InfoUseCase {
	return InfoUseCase{
		formQ:  rgst.Repository().NewFormQuery(),
		access: rgst.Service().AccessController(),
	}
}

func (uc InfoUseCase) Do(ipt InfoInput) (*InfoOutput, error) {
	if !uc.access.CanAccessOrg(ipt.Ctx, ipt.UserID, ipt.OrgID) {
		return nil, exception.ErrUnauthorized
	}
	f, err := uc.formQ.GetByID(ipt.Ctx, ipt.FormID)
	if err != nil {
		return nil, err
	}
	if !f.IsOpen {
		return nil, exception.ErrNotAvailable
	}
	return &InfoOutput{
		Form: *f,
	}, nil
}
