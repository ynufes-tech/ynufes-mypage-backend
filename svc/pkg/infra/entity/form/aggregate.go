package entity

import (
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

const FormCollectionName = "Forms"

type (
	Form struct {
		ID          string  `firestore:"-"`
		EventID     int64   `firestore:"event_id"`
		Title       string  `firestore:"title"`
		Summary     string  `firestore:"summary"`
		Description string  `firestore:"description"`
		Roles       []int64 `firestore:"roles"`
		Deadline    int64   `firestore:"deadline"`
		IsOpen      bool    `firestore:"is_open"`
	}
)

func (f Form) ToModel() (*form.Form, error) {
	fid, err := identity.ImportID(f.ID)
	if err != nil {
		return nil, err
	}
	eID := identity.NewID(f.EventID)
	roles := make([]user.RoleID, len(f.Roles))
	for i, r := range f.Roles {
		roles[i] = identity.NewID(r)
	}
	deadline := time.UnixMilli(f.Deadline)
	return form.NewForm(
		fid, eID, f.Title, f.Summary, f.Description, roles, deadline, f.IsOpen,
	), nil
}
