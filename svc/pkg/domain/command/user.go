package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type User interface {
	Create(context.Context, user.User) error
	UpdateLine(ctx context.Context, oldUser *user.User, update user.Line) error
	UpdateUserDetail(ctx context.Context, oldUser *user.User, update user.Detail) error
	UpdateAgent(ctx context.Context, oldUser *user.User, update user.Agent) error
	UpdateAdmin(ctx context.Context, oldUser *user.User, update user.Admin) error
	UpdateAll(context.Context, user.User) error
	Delete(context.Context, user.User) error
}
