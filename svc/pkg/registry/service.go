package registry

import (
	linePkg "ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/svc/pkg/domain/service/access"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
)

type Service struct{}

var (
	lineAuthVerifier lineDomain.AuthVerifier
	accessController access.AccessController
)

func init() {
	config := setting.Get()

	lineAuthVerifier = linePkg.NewAuthVerifier(
		config.ThirdParty.LineLogin.CallbackURI,
		config.ThirdParty.LineLogin.ClientID,
		config.ThirdParty.LineLogin.ClientSecret,
	)
	accessController = access.NewAccessController(repo.NewRelationQuery())
}

func NewService() Service {
	return Service{}
}

func (s Service) LineAuthVerifier() *lineDomain.AuthVerifier {
	return &lineAuthVerifier
}

func (s Service) AccessController() *access.AccessController {
	return &accessController
}
