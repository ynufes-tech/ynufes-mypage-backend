package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

type Question interface {
	GetByID(ctx context.Context, id id.QuestionID) (*question.Question, error)
	ListByFormID(ctx context.Context, id id.FormID) ([]question.Question, error)
}
