package form

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	// Form TODO: Implement Sections field
	Form struct {
		ID          ID
		EventID     event.ID
		Title       string
		Summary     string
		Description string
		Roles       []user.RoleID
		//Sections    []SectionID
	}

	ID util.ID
	//SectionID util.ID
)

func NewForm(id ID, eventID event.ID, title, summary, description string, roles []user.RoleID) *Form {
	return &Form{
		ID:          id,
		EventID:     eventID,
		Title:       title,
		Summary:     summary,
		Description: description,
		Roles:       roles,
	}
}
