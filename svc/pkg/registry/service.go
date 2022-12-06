package registry

import (
	linePkg "ynufes-mypage-backend/pkg/line"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
)

func NewAuthStateManager() lineDomain.AuthStateManager {
	return linePkg.NewAuthStateManager()
}
