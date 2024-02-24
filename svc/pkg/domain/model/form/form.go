package form

import (
	"sort"
	"time"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	Form struct {
		ID          id.FormID
		EventID     id.EventID
		Title       string
		Summary     string
		Description string
		Roles       []id.RoleID
		Deadline    time.Time
		IsOpen      bool
		Sections    SectionsOrder
	}
	SectionsOrder map[id.SectionID]float64
)

func NewForm(
	id id.FormID,
	eventID id.EventID,
	title, summary, description string,
	sectionOrders map[id.SectionID]float64,
	roles []id.RoleID,
	deadline time.Time,
	isOpen bool,
) *Form {
	return &Form{
		ID:          id,
		EventID:     eventID,
		Title:       title,
		Summary:     summary,
		Description: description,
		Roles:       roles,
		Deadline:    deadline,
		IsOpen:      isOpen,
		Sections:    sectionOrders,
	}
}

func (o SectionsOrder) GetOrderedIDs() []id.SectionID {
	ids := make([]id.SectionID, 0, len(o))
	for oid := range o {
		ids = append(ids, oid)
	}
	sort.Slice(ids, func(i, j int) bool {
		return o[ids[i]] < o[ids[j]]
	})
	return ids
}
