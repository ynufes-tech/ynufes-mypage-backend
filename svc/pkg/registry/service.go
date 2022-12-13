package registry

import (
	linePkg "ynufes-mypage-backend/pkg/line"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (Service) NewAuthStateManager() lineDomain.AuthStateManager {
	return linePkg.NewAuthStateManager()
}
