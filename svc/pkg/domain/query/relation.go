package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type Relation interface {
	ListUserIDsByOrgID(ctx context.Context, orgID id.OrgID) ([]id.UserID, error)
	ListOrgIDsByUserID(ctx context.Context, userID id.UserID) (id.OrgIDs, error)
}
