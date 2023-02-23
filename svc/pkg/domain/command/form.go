package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
)

type Form interface {
	Create(context.Context, form.Form) error
	Set(context.Context, form.Form) error
}
