package registry

import (
	"ynufes-mypage-backend/pkg/identity"
	linePkg "ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/pkg/setting"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
	utilSVC "ynufes-mypage-backend/svc/pkg/domain/service/util"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) NewLineAuthVerifier() lineDomain.AuthVerifier {
	config := setting.Get()
	return linePkg.NewAuthVerifier(config.ThirdParty.LineLogin.CallbackURI, config.ThirdParty.LineLogin.ClientID, config.ThirdParty.LineLogin.ClientSecret)
}

func (s Service) NewIDManager() utilSVC.IDManager {
	return identity.NewIDManager()
}
