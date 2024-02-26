package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
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

// Set sets the staff entity, irrelevant of the existence of the entity.
// Note that this does not check if the user exists.
func (s Staff) Set(ctx context.Context, stf staff.Staff) error {
	if stf.UserID == nil || !stf.UserID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	return s.ref.Child(stf.UserID.ExportID()).
		Set(ctx, entity.Staff{
			IsAdmin: stf.IsAdmin,
		})
}
