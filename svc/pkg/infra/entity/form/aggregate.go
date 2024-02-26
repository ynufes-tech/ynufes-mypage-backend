package entity

import (
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
)

const FormRootName = "Forms"

type (
	Form struct {
		ID          id.FormID          `json:"-"`
		EventID     string             `json:"event_id"`
		Title       string             `json:"title"`
		Summary     string             `json:"summary"`
		Description string             `json:"description"`
		Sections    map[string]float64 `json:"sections"`
		Roles       map[string]bool    `json:"roles"`
		Deadline    int64              `json:"deadline"`
		IsOpen      bool               `json:"is_open"`
	}
)

func NewForm(
	id id.FormID, eventID string,
	title, summary, description string,
	sections map[string]float64,
	roles map[string]bool,
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
	if f.EventID == "" {
		return nil, exception.ErrIDNotAssigned
	}

	eID, err := identity.ImportID(f.EventID)

	if err != nil {
		return nil, err
	}

	roles := make([]id.RoleID, 0, len(f.Roles))
	for k, v := range f.Roles {
		if !v {
			continue
		}
		tid, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		roles = append(roles, tid)
	}

	sections := make(map[id.SectionID]float64, len(f.Sections))
	for k, v := range f.Sections {
		tid, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		sections[tid] = v
	}

	deadline := time.UnixMilli(f.Deadline)
	return form.NewForm(
		f.ID, eID,
		f.Title, f.Summary, f.Description,
		sections, roles, deadline, f.IsOpen,
	), nil
}
