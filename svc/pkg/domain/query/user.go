package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type User interface {
	GetByID(context.Context, user.ID) (*user.User, error)
	GetByLineServiceID(context.Context, user.LineServiceID) (*user.User, error)
}
