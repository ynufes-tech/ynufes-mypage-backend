package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
)

type (
	Form struct {
		ID          string              `firestore:"-"`
		Title       string              `firestore:"title"`
		Summary     string              `firestore:"summary"`
		Description string              `firestore:"description"`
		Questions   map[string]Question `firestore:"questions"`
	}
	Question struct {
		ID           string                 `firestore:"-"`
		QuestionText string                 `firestore:"text"`
		QuestionType int                    `firestore:"type"`
		Properties   map[string]interface{} `firestore:"props"`
		Order        int                    `firestore:"order"`
	}
)

func (f Form) ToModel() (*form.Form, error) {
	fid, err := identity.ImportID(f.ID)
	if err != nil {
		return nil, err
	}
	var questions []form.Question
	for _, v := range f.Questions {
		q, err := ImportQuestion(v)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return &form.Form{
		ID:          fid,
		Title:       f.Title,
		Summary:     f.Summary,
		Description: f.Description,
		Questions:   questions,
	}, nil
}
