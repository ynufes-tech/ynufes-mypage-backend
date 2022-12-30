package line

import (
	"context"
	"log"
	linePkg "ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/domain/service/line"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type AuthUseCase struct {
	authVerifier line.AuthVerifier
	userQ        query.User
	userC        command.User
}

type AuthInput struct {
	State string
	Code  string
	Ctx   context.Context
}

type AuthOutput struct {
	AccessToken  string
	RefreshToken string
	ErrorMsg     string
	UserInfo     *user.User
}

// TODO: handler.goの内容を分割する
func NewAuthCodeUseCase(rgst registry.Registry) AuthUseCase {
	return AuthUseCase{
		authVerifier: rgst.Service().NewLineAuthVerifier(),
		userQ:        rgst.Repository().NewUserQuery(),
		userC:        rgst.Repository().NewUserCommand(),
	}
}

func (uc AuthUseCase) Do(ipt AuthInput) (*AuthOutput, error) {
	token, err := uc.authVerifier.RequestAccessToken(ipt.Code, ipt.State)
	if err != nil {
		log.Println("Failed to get access token from LINE server... ", err)
		return nil, err
	}
	profile, err := linePkg.GetProfile(token.AccessToken)
	if err != nil {
		// failed to get profile
		log.Printf("failed to get profile: %v", err)
		return nil, err
	}
	encryptedAccessToken, err := user.NewEncryptedAccessToken(user.PlainAccessToken(token.AccessToken))
	if err != nil {
		return nil, err
	}
	encryptedRefreshToken, err := user.NewEncryptedRefreshToken(user.PlainRefreshToken(token.RefreshToken))
	if err != nil {
		return nil, err
	}
	return &AuthOutput{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Code:         CodeSuccess,
		LineInfo: user.Line{
			LineServiceID:         user.LineServiceID(profile.UserID),
			LineProfilePictureURL: user.LineProfilePictureURL(profile.PictureURL),
			LineDisplayName:       profile.DisplayName,
			EncryptedAccessToken:  encryptedAccessToken,
			EncryptedRefreshToken: encryptedRefreshToken,
		},
	}, nil
}
