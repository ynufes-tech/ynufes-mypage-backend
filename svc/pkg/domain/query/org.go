package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type Org interface {
	GetByID(ctx context.Context, id org.ID) (*org.Org, error)
	ListByGrantedUserID(ctx context.Context, id user.ID) ([]org.Org, error)
}
