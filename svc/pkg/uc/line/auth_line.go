package line

import (
	"ynufes-mypage-backend/svc/pkg/domain/service/line"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type AuthUseCase struct {
	authVerifier line.AuthVerifier
}

type AuthInput struct {
	State string
	Code  string
}

type AuthOutput struct {
	AccessToken  string
	RefreshToken string
}

func NewAuthCodeUseCase(rgst registry.Registry) AuthUseCase {
	return AuthUseCase{
		authVerifier: rgst.Service().NewLineAuthVerifier(),
	}
}

func (u AuthUseCase) Do(ipt AuthInput) (*AuthOutput, error) {
	result, err := u.authVerifier.RequestAccessToken(ipt.Code, ipt.State)
	if err != nil {
		return nil, err
	}
	return &AuthOutput{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}, nil
}
