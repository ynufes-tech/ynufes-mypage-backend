package form

import (
	"time"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type (
	Form struct {
		ID          id.FormID
		EventID     id.EventID
		Title       string
		Summary     string
		Description string
		Roles       []user.RoleID
		Deadline    time.Time
		IsOpen      bool
		Sections    []id.SectionID
	}
)

func NewForm(
	id id.FormID,
	eventID id.EventID,
	title, summary, description string,
	sectionIDs []id.SectionID,
	roles []user.RoleID,
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
		Sections:    sectionIDs,
	}
}
