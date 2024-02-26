package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/relation"
)

type RelationOrgUser struct {
	orgUserRef *db.Ref
}

func NewRelationOrgUser(db *firebase.Firebase) RelationOrgUser {
	return RelationOrgUser{
		orgUserRef: db.Client(entity.RelationRootName).Child(entity.RelationOrgUserName),
	}
}

func (r RelationOrgUser) ListUserIDsByOrgID(ctx context.Context, orgID id.OrgID) ([]id.UserID, error) {
	qs, err := r.orgUserRef.OrderByChild("org_id").
		EqualTo(orgID.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}

	userIDs := make([]id.UserID, len(qs))
	for i := range qs {
		var rEntity entity.OrgUserRelation
		if err := qs[i].Unmarshal(&rEntity); err != nil {
			return nil, err
		}
		uid, err := identity.ImportID(rEntity.UserID)
		if err != nil {
			return nil, err
		}
		userIDs[i] = uid
	}
	return userIDs, nil
}

func (r RelationOrgUser) ListOrgIDsByUserID(ctx context.Context, userID id.UserID) (id.OrgIDs, error) {
	qs, err := r.orgUserRef.OrderByChild("user_id").
		EqualTo(userID.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}

	orgIDs := make([]id.OrgID, len(qs))
	for i := range qs {
		var rEntity entity.OrgUserRelation
		if err := qs[i].Unmarshal(&rEntity); err != nil {
			return nil, err
		}
		oid, err := identity.ImportID(rEntity.OrgID)
		if err != nil {
			return nil, err
		}
		orgIDs[i] = oid
	}
	return orgIDs, nil
}
