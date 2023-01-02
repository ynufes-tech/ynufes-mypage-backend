package line

import (
	"context"
	"fmt"
	"log"
	linePkg "ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/pkg/snowflake"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/domain/service/line"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type AuthUseCase struct {
	authVerifier *line.AuthVerifier
	userQ        query.User
	userC        command.User
	enableLine   bool
}

type AuthInput struct {
	State string
	Code  string
	Ctx   context.Context
}

type AuthOutput struct {
	ErrorMsg string
	UserInfo *user.User
}

func NewAuthCodeUseCase(rgst registry.Registry, enableLineAuth bool, authVerifier *line.AuthVerifier) AuthUseCase {
	return AuthUseCase{
		authVerifier: authVerifier,
		userQ:        rgst.Repository().NewUserQuery(),
		userC:        rgst.Repository().NewUserCommand(),
		enableLine:   enableLineAuth,
	}
}

func (uc AuthUseCase) Do(ipt AuthInput) (*AuthOutput, error) {
	var aToken user.EncryptedAccessToken
	var rToken user.EncryptedRefreshToken
	var profile linePkg.ProfileResponse
	if uc.enableLine {
		token, err := (*uc.authVerifier).RequestAccessToken(ipt.Code, ipt.State)
		if err != nil {
			err = fmt.Errorf("bad request, failed to authorize with LINE: %v", err)
			log.Printf("error: %v", err)
			return &AuthOutput{
				ErrorMsg: err.Error(),
			}, nil
		}
		aToken = user.NewEncryptedAccessToken(user.PlainAccessToken(token.AccessToken))
		rToken = user.NewEncryptedRefreshToken(user.PlainRefreshToken(token.RefreshToken))
		profile, err = linePkg.GetProfile(token.AccessToken)
		if err != nil {
			// failed to get profile
			log.Printf("failed to get profile: %v", err)
			return nil, err
		}
	} else {
		aToken = user.NewEncryptedAccessToken("testAccessToken")
		rToken = user.NewEncryptedRefreshToken("testRefreshToken")
		profile = linePkg.ProfileResponse{
			UserID:        "testUserLineID",
			DisplayName:   "testUserDisplayName",
			PictureURL:    "https://testUserPicture.com",
			StatusMessage: "testUserStatusMessage",
		}
	}
	lineServiceID := user.LineServiceID(profile.UserID)
	u, err := uc.userQ.GetByLineServiceID(ipt.Ctx, lineServiceID)
	if err != nil {
		// if error is "user not found", Create User and redirect to basic info form
		// Otherwise, respond with error
		newID := user.ID(snowflake.NewSnowflake())
		newUser := user.User{
			ID:     newID,
			Status: user.StatusNew,
			Line: user.Line{
				LineServiceID:         lineServiceID,
				LineProfilePictureURL: user.LineProfilePictureURL(profile.PictureURL),
				LineDisplayName:       profile.DisplayName,
				EncryptedAccessToken:  aToken,
				EncryptedRefreshToken: rToken,
			},
			Detail: user.Detail{
				Name:      user.Name{},
				Email:     "",
				Gender:    user.GenderNotSpecified,
				StudentID: "",
				Type:      user.TypeNormal,
			},
		}
		if err = uc.userC.Create(ipt.Ctx, newUser); err != nil {
			log.Printf("failed to create user: %v", err)
			return nil, err
		}
		return &AuthOutput{
			UserInfo: &newUser,
		}, nil
	}
	// User found. Update Line info
	update := user.Line{
		LineServiceID:         lineServiceID,
		LineProfilePictureURL: user.LineProfilePictureURL(profile.PictureURL),
		LineDisplayName:       profile.DisplayName,
		EncryptedAccessToken:  aToken,
		EncryptedRefreshToken: rToken,
	}
	if err := uc.userC.UpdateLine(ipt.Ctx, u, update); err != nil {
		return nil, fmt.Errorf("failed to update line info: %v", err)
	}
	return &AuthOutput{
		UserInfo: u,
	}, nil
}
