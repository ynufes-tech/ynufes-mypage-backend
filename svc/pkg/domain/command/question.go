package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

type Question interface {
	Create(context.Context, *question.Question) error
	Set(context.Context, question.Question) error
	UpdateCustoms(context.Context, map[string]interface{}) error
}
