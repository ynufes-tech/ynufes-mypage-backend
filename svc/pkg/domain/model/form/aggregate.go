package form

import "ynufes-mypage-backend/svc/pkg/domain/model/util"

type (
	// Form TODO: Implement Sections field
	Form struct {
		ID          ID
		Title       string
		Summary     string
		Description string
		//Sections    []SectionID
	}

	ID util.ID
	//SectionID util.ID
)

func NewForm(id ID, title, summary, description string) *Form {
	return &Form{
		ID:          id,
		Title:       title,
		Summary:     summary,
		Description: description,
	}
}
