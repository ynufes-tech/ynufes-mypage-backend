package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type Relation interface {
	CreateOrgUser(ctx context.Context, orgID id.OrgID, userID id.UserID) error
	DeleteOrgUser(ctx context.Context, orgID id.OrgID, userID id.UserID) error
}
