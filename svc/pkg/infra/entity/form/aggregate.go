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
		ID           string    `firestore:"-"`
		EventID      int64     `firestore:"event_id"`
		Title        string    `firestore:"title"`
		Summary      string    `firestore:"summary"`
		Description  string    `firestore:"description"`
		Roles        []int64   `firestore:"roles"`
		Deadline     int64     `firestore:"deadline"`
		IsOpen       bool      `firestore:"is_open"`
		SectionOrder []int64   `firestore:"section_order"`
		Sections     []Section `firestore:"sections"`
	}
)

func NewForm(
	id string, eventID int64, title, summary, description string, roles []int64, deadline int64, isOpen bool,
	sectionOrder []int64, sections []Section,
) Form {
	return Form{
		ID:           id,
		EventID:      eventID,
		Title:        title,
		Summary:      summary,
		Description:  description,
		Roles:        roles,
		Deadline:     deadline,
		IsOpen:       isOpen,
		SectionOrder: sectionOrder,
		Sections:     sections,
	}
}

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
	sectionsOrder := make([]form.SectionID, len(f.SectionOrder))
	for i := range sectionsOrder {
		sectionsOrder[i] = identity.NewID(f.SectionOrder[i])
	}

	sections := make(map[form.SectionID]form.Section, len(f.Sections))
	for _, s := range f.Sections {
		m, err := s.ToModel()
		if err != nil {
			return nil, err
		}
		sections[identity.NewID(s.ID)] = *m
	}

	deadline := time.UnixMilli(f.Deadline)
	return form.NewForm(
		fid, eID, f.Title, f.Summary, f.Description, roles, deadline, f.IsOpen,
		sectionsOrder, sections,
	), nil
}
