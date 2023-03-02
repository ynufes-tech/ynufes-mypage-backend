package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type User interface {
	Create(context.Context, *user.User) error
	Set(context.Context, user.User) error
	UpdateLine(context.Context, id.UserID, user.Line) error
	UpdateUserDetail(context.Context, id.UserID, user.Detail) error
	UpdateAgent(context.Context, id.UserID, user.Agent) error
	UpdateAdmin(context.Context, id.UserID, user.Admin) error
	Delete(context.Context, user.User) error
}
