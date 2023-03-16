package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/line"
)

type Line interface {
	Create(context.Context, line.LineUser) error
	Set(context.Context, line.LineUser) error
}
