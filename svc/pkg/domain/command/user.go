package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type User interface {
	Create(context.Context, user.User) error
	UpdateLine(context.Context, user.User) error
	UpdateAll(context.Context, user.User) error
	Delete(context.Context, user.User) error
}
