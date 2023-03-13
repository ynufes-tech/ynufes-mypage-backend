package section

import (
	"ynufes-mypage-backend/svc/pkg/registry"
	"ynufes-mypage-backend/svc/pkg/uc/section"
)

type Section struct {
	infoUC section.InfoUseCase
}

func NewSection(rgst registry.Registry) Section {
	return Section{
		infoUC: section.NewInfo(rgst),
	}
}
