package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/relation"
)

type RoleStaffRelation struct {
	ref *db.Ref
}

func NewRelationRoleStaff(c *firebase.Firebase) RoleStaffRelation {
	return RoleStaffRelation{
		ref: c.Client(entity.RelationRootName).Child(entity.RelationRoleStaffName),
	}
}

func (r RoleStaffRelation) CreateRoleStaff(ctx context.Context, roleID id.RoleID, staffID id.UserID) error {
	t := entity.RoleStaffRelation{
		RoleID: roleID.ExportID(),
		UserID: staffID.ExportID(),
	}
	_, err := r.ref.
		Push(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (r RoleStaffRelation) DeleteRoleStaff(ctx context.Context, roleID id.RoleID, staffID id.UserID) error {
	relations, err := r.ref.
		OrderByChild("user_id").EqualTo(staffID.ExportID()).
		GetOrdered(ctx)
	if err != nil {
		return err
	}
	var found bool
	for _, relation := range relations {
		var rEntity entity.RoleStaffRelation
		if err := relation.Unmarshal(&rEntity); err != nil {
			return err
		}
		if rEntity.RoleID != roleID.ExportID() {
			continue
		}
		if err := r.ref.Child(relation.Key()).Delete(ctx); err != nil {
			return err
		}
		found = true
	}
	if !found {
		return exception.ErrNotFound
	}
	return nil
}
