package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
)

type Org interface {
	GetByID(ctx context.Context, id id.OrgID) (*org.Org, error)
}
