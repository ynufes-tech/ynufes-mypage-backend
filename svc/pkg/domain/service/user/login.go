package user

import (
	"context"
	"ynufes-mypage-backend/pkg/jwt"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type Login interface {
	Do(ctx context.Context, jwt jwt.JWT) (user.ID, error)
}
