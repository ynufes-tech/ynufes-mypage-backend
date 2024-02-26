package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/staff"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/staff"
)

type Staff struct {
	ref *db.Ref
}

func NewStaff(c *firebase.Firebase) *Staff {
	return &Staff{
		ref: c.Client(entity.StaffTableName),
	}
}

func (s Staff) GetStaffByUserID(ctx context.Context, uid id.UserID) (*staff.Staff, error) {
	if uid == nil || !uid.HasValue() {
		return nil, exception.ErrIDNotAssigned
	}
	r, err := s.ref.OrderByKey().EqualTo(uid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	if len(r) > 1 {
		return nil, fmt.Errorf("multiple staff found with id: %s", uid)
	}
	var staffEntity entity.Staff
	if err := r[0].Unmarshal(&staffEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal staff entity: %w", err)
	}
	return &staff.Staff{
		UserID:  uid,
		IsAdmin: staffEntity.IsAdmin,
	}, nil
}
