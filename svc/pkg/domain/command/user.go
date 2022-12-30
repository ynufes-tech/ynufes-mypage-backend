package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type User interface {
	Create(context.Context, user.User) error
	UpdateLine(ctx context.Context, oldUser *user.User, update user.Line) error
	UpdateAll(context.Context, user.User) error
	Delete(context.Context, user.User) error
}
