package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/relation"
)

type RelationOrgUser struct {
	OrgUserRef *db.Ref
}

func NewRelationOrgUser(db *firebase.Firebase) RelationOrgUser {
	return RelationOrgUser{
		OrgUserRef: db.Client(entity.RelationRootName).Child(entity.RelationOrgUserName),
	}
}

func (r RelationOrgUser) CreateOrgUser(ctx context.Context, orgID id.OrgID, userID id.UserID) error {
	t := entity.OrgUserRelation{
		OrgID:  orgID.ExportID(),
		UserID: userID.ExportID(),
	}
	_, err := r.OrgUserRef.
		Push(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (r RelationOrgUser) DeleteOrgUser(ctx context.Context, orgID id.OrgID, userID id.UserID) error {
	orgs, err := r.OrgUserRef.
		OrderByChild("user_id").EqualTo(userID.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return err
	}
	var found bool
	for _, org := range orgs {
		var rEntity entity.OrgUserRelation
		if err := org.Unmarshal(&rEntity); err != nil {
			return err
		}
		fmt.Println(rEntity)
		if rEntity.OrgID != orgID.ExportID() {
			continue
		}
		if err := r.OrgUserRef.Child(org.Key()).Delete(ctx); err != nil {
			return err
		}
		found = true
	}
	if !found {
		return exception.ErrNotFound
	}
	return nil
}
