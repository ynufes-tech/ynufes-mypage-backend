package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

type Question interface {
	GetByID(ctx context.Context, id question.ID) (*question.Question, error)
	ListByEventID(ctx context.Context, id event.ID) ([]question.Question, error)
	ListByFormID(ctx context.Context, id form.ID) ([]question.Question, error)
}
