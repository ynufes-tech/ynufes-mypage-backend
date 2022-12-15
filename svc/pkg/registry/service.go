package registry

import (
	linePkg "ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/pkg/setting"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) NewLineAuthVerifier() lineDomain.AuthVerifier {
	config := setting.Get()
	return linePkg.NewAuthVerifier(config.ThirdParty.LineLogin.CallbackURI, config.ThirdParty.LineLogin.ClientID, config.ThirdParty.LineLogin.ClientSecret)
}
