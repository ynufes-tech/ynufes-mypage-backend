package entity

import (
	"fmt"
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
		ID           string      `firestore:"question_id"`
		QuestionText string      `firestore:"question"`
		QuestionType int         `firestore:"question_type"`
		Properties   interface{} `firestore:"properties"`
		Order        int         `firestore:"order"`
	}
)

func (f Form) ToModel() (*form.Form, error) {
	fid, err := identity.ImportID(f.ID)
	if err != nil {
		return nil, err
	}
	questions := make(map[form.QID]form.Question)
	for _, v := range f.Questions {
		qid, err := identity.ImportID(v.ID)
		if err != nil {
			return nil, err
		}
		properties, err := importQuestionProperties(form.QuestionType(v.QuestionType), v.Properties)
		if err != nil {
			return nil, fmt.Errorf("failed to import question properties in ToModel(): %w", err)
		}
		questions[qid] = form.Question{
			ID:           qid,
			Type:         form.QuestionType(v.QuestionType),
			QuestionText: v.QuestionText,
			Order:        0,
			Properties:   *properties,
		}
	}
	return &form.Form{
		ID:          fid,
		Title:       f.Title,
		Summary:     f.Summary,
		Description: f.Description,
		Questions:   questions,
	}, nil
}
