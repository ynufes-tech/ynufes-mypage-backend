package entity

import (
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

const FormRootName = "Forms"

type (
	Form struct {
		ID          id.FormID `json:"-"`
		EventID     int64     `json:"event_id"`
		Title       string    `json:"title"`
		Summary     string    `json:"summary"`
		Description string    `json:"description"`
		Roles       []int64   `json:"roles"`
		Deadline    int64     `json:"deadline"`
		IsOpen      bool      `json:"is_open"`
		Sections    []int64   `json:"section"`
	}
)

func NewForm(
	id id.FormID, eventID int64,
	title, summary, description string,
	sections, roles []int64,
	deadline int64,
	isOpen bool,
) Form {
	return Form{
		ID:          id,
		EventID:     eventID,
		Title:       title,
		Summary:     summary,
		Description: description,
		Roles:       roles,
		Deadline:    deadline,
		IsOpen:      isOpen,
		Sections:    sections,
	}
}

func (f Form) ToModel() (*form.Form, error) {
	eID := identity.NewID(f.EventID)
	roles := make([]user.RoleID, len(f.Roles))
	for i, r := range f.Roles {
		roles[i] = identity.NewID(r)
	}
	sections := make([]id.SectionID, len(f.Sections))
	for i := range sections {
		sections[i] = identity.NewID(f.Sections[i])
	}

	deadline := time.UnixMilli(f.Deadline)
	return form.NewForm(
		f.ID, eID,
		f.Title, f.Summary, f.Description,
		sections, roles, deadline, f.IsOpen,
	), nil
}
