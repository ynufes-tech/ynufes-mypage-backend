package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

type Question interface {
	Create(ctx context.Context, q question.Question) error
	UpdateCustoms(ctx context.Context, id question.ID, customs map[string]interface{}) error
	Set(ctx context.Context, q question.Question) error
}
