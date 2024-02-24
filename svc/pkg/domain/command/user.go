package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type User interface {
	Create(context.Context, *user.User) error
	Set(context.Context, user.User) error
	UpdateUserDetail(context.Context, id.UserID, user.Detail) error
	SetAgent(context.Context, id.UserID, user.Agent) error
	Delete(context.Context, user.User) error
}
