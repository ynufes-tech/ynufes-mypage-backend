package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/line"
)

type Line interface {
	GetByUserID(context.Context, id.UserID) (*line.LineUser, error)
	GetByLineServiceID(context.Context, line.LineServiceID) (*line.LineUser, error)
}
