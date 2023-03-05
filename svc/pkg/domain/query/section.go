package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
)

type Section interface {
	ListSectionsByFormID(ctx context.Context, fid id.FormID) ([]section.Section, error)
	GetSectionByID(ctx context.Context, tid id.SectionID) (*section.Section, error)
}
