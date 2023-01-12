package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
)

type Org interface {
	GetByID(ctx context.Context, id org.ID) (*org.Org, error)
}
