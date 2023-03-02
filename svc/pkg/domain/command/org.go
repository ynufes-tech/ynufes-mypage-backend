package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
)

type Org interface {
	Create(context.Context, *org.Org) error
	Set(context.Context, org.Org) error
	UpdateUsers(context.Context, []id.UserID) error
	UpdateIsOpen(context.Context, bool) error
}
