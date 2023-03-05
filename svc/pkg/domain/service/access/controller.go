package access

import (
	"context"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/query"
)

type AccessController struct {
	relationQ query.Relation
}

func NewAccessController(relationQ query.Relation) AccessController {
	return AccessController{
		relationQ: relationQ,
	}
}

func (c AccessController) CanAccessOrg(ctx context.Context, uid id.UserID, oid id.OrgID) bool {
	orgs, err := c.relationQ.ListOrgIDsByUserID(ctx, uid)
	if err != nil {
		fmt.Printf("failed to get orgs in CanAccessOrg: %v\n", err)
		return false
	}
	for _, tid := range orgs {
		if tid == oid {
			return true
		}
	}
	return false
}
