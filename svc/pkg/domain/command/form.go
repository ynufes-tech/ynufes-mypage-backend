package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type Form interface {
	Create(context.Context, *form.Form) error
	Set(context.Context, form.Form) error
	AddSectionOrder(context.Context, id.FormID, id.SectionID, float64) error
	UpdateSectionOrder(context.Context, id.FormID, id.SectionID, float64) error
}
