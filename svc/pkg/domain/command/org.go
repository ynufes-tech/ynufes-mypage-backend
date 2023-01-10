package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
)

type Org interface {
	Create(context.Context, *org.Org) error
	UpdateAll(context.Context, *org.Org) error
	Delete(context.Context, *org.Org) error
}
