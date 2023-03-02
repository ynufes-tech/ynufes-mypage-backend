package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

type Question interface {
	Create(ctx context.Context, f id.FormID, q question.Question) error
	UpdateCustoms(ctx context.Context, id id.QuestionID, customs map[string]interface{}) error
	Set(ctx context.Context, f id.FormID, q question.Question) error
}
