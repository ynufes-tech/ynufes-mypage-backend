package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
)

const FormCollectionName = "Forms"

type (
	Form struct {
		ID          string `firestore:"-"`
		Title       string `firestore:"title"`
		Summary     string `firestore:"summary"`
		Description string `firestore:"description"`
	}
)

func (f Form) ToModel() (*form.Form, error) {
	fid, err := identity.ImportID(f.ID)
	if err != nil {
		return nil, err
	}
	return form.NewForm(fid, f.Title, f.Summary, f.Description), nil
}
