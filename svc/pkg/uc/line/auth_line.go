package line

import (
	"context"
	"log"
	line2 "ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/domain/service/line"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type ResultCode int

const (
	CodeSuccess ResultCode = 200
	CodeFailed  ResultCode = 400
	CodeError   ResultCode = 500
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
	Code         ResultCode
	ErrorMsg     string
	LineInfo     user.Line
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
	profile, err := line2.GetProfile(token.AccessToken)
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
