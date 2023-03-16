package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type User interface {
	GetByID(context.Context, id.UserID) (*user.User, error)
}
