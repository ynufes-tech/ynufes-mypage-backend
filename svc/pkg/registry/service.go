package registry

import (
	"time"
	linePkg "ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/pkg/token"
	"ynufes-mypage-backend/svc/pkg/domain/service/access"
	"ynufes-mypage-backend/svc/pkg/domain/service/auth"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
)

type Service struct{}

var (
	lineAuthVerifier lineDomain.AuthVerifier
	accessController access.AccessController
	tokenIssuer      auth.TokenIssuer
)

func init() {
	config := setting.Get()

	lineAuthVerifier = linePkg.NewAuthVerifier(
		config.ThirdParty.LineLogin.CallbackURI,
		config.ThirdParty.LineLogin.ClientID,
		config.ThirdParty.LineLogin.ClientSecret,
	)
	accessController = access.NewAccessController(repo.NewRelationQuery())
	tokenIssuer = token.NewTokenIssuer(
		config.Application.Authentication.JwtSecret,
		config.Application.Server.Backend.Domain,
		24*time.Hour,
	)
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

func (s Service) TokenIssuer() *auth.TokenIssuer {
	return &tokenIssuer
}
