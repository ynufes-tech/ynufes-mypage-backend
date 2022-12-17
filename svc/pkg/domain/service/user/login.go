package user

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type Login interface {
	Do(ctx context.Context, jwt user.JWT) (user.ID, error)
}
