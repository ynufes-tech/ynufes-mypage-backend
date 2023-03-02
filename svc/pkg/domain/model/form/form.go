package form

import (
	"time"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	Form struct {
		ID          ID
		EventID     event.ID
		Title       string
		Summary     string
		Description string
		Roles       []user.RoleID
		Deadline    time.Time
		IsOpen      bool
		Sections    []section.ID
	}

	ID util.ID
)

func NewForm(
	id ID,
	eventID event.ID,
	title, summary, description string,
	sectionIDs []section.ID,
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
