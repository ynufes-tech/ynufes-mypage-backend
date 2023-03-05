package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
)

type Section interface {
	Create(ctx context.Context, targetSection *section.Section) error
	Set(ctx context.Context, targetSection section.Section) error
	LinkQuestion(ctx context.Context, secID id.SectionID, qID id.QuestionID, index float64) error
	UnlinkQuestion(ctx context.Context, secID id.SectionID, qID id.QuestionID) error
}
