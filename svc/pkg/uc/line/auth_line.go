package line

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	linePkg "ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/line"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	lineSVC "ynufes-mypage-backend/svc/pkg/domain/service/line"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type AuthUseCase struct {
	authVerifier *lineSVC.AuthVerifier
	userQ        query.User
	userC        command.User
	lineQ        query.Line
	lineC        command.Line
	enableLine   bool
}

type AuthInput struct {
	State string
	Code  string
	Ctx   context.Context
}

type AuthOutput struct {
	UserInfo      *user.User
	LineServiceID line.LineServiceID
}

func NewAuthCodeUseCase(rgst registry.Registry, enableLineAuth bool, authVerifier *lineSVC.AuthVerifier) AuthUseCase {
	return AuthUseCase{
		authVerifier: authVerifier,
		userQ:        rgst.Repository().NewUserQuery(),
		userC:        rgst.Repository().NewUserCommand(),
		lineQ:        rgst.Repository().NewLineQuery(),
		lineC:        rgst.Repository().NewLineCommand(),
		enableLine:   enableLineAuth,
	}
}

func (uc AuthUseCase) Do(ipt AuthInput) (*AuthOutput, error) {
	var aToken line.EncryptedAccessToken
	var rToken line.EncryptedRefreshToken
	var profile linePkg.ProfileResponse
	if uc.enableLine {
		token, err := (*uc.authVerifier).RequestAccessToken(ipt.Code, ipt.State)
		if err != nil {
			err = fmt.Errorf("bad request, failed to authorize with LINE: %v", err)
			log.Printf("error in line auth: %v\n", err)
			return nil, err
		}
		aToken = line.NewEncryptedAccessToken(line.PlainAccessToken(token.AccessToken))
		rToken = line.NewEncryptedRefreshToken(line.PlainRefreshToken(token.RefreshToken))
		profile, err = linePkg.GetProfile(token.AccessToken)
		if err != nil {
			// failed to get profile
			err = fmt.Errorf("bad request, failed to get profile from LINE: %v", err)
			log.Printf("error in line auth: %v\n", err)
			return nil, err
		}
	} else {
		// if line auth is disabled, return dummy data
		// if the request has query, it will be used.
		c := ipt.Ctx.(*gin.Context)
		aToken = line.NewEncryptedAccessToken(
			line.PlainAccessToken(c.DefaultQuery("accessToken", "testAccessToken")))
		rToken = line.NewEncryptedRefreshToken(
			line.PlainRefreshToken(c.DefaultQuery("refreshToken", "testRefreshToken")))
		profile = linePkg.ProfileResponse{
			UserID:        c.DefaultQuery("userID", "testUserID"),
			DisplayName:   c.DefaultQuery("displayName", "testDisplayName"),
			PictureURL:    c.DefaultQuery("pictureURL", "https://testUserPicture.com"),
			StatusMessage: c.DefaultQuery("statusMessage", "testStatusMessage"),
		}
	}
	lineServiceID := line.LineServiceID(profile.UserID)

	lu, err := uc.lineQ.GetByLineServiceID(ipt.Ctx, lineServiceID)

	if err != nil {
		// if error is "user not found", Create User and redirect to basic info form
		// Otherwise, respond with error
		newUser := user.User{
			Detail: user.Detail{
				Name:       user.Name{},
				Email:      "",
				Gender:     user.GenderUnknown,
				StudentID:  "",
				Type:       user.TypeNormal,
				PictureURL: user.PictureURL(profile.PictureURL),
			},
			Admin: user.Admin{},
			Agent: user.Agent{},
		}
		if err := uc.userC.Create(ipt.Ctx, &newUser); err != nil {
			err = fmt.Errorf("failed to create user: %v", err)
			log.Printf("error in line auth: %v\n", err)
			return nil, err
		}
		if err := uc.lineC.Create(ipt.Ctx,
			line.LineUser{
				UserID:                newUser.ID,
				LineServiceID:         lineServiceID,
				LineDisplayName:       profile.DisplayName,
				EncryptedAccessToken:  aToken,
				EncryptedRefreshToken: rToken,
			}); err != nil {
			err = fmt.Errorf("failed to create line user: %v", err)
			log.Printf("error in line auth: %v\n", err)
			return nil, err
		}
		return &AuthOutput{
			UserInfo:      &newUser,
			LineServiceID: lineServiceID,
		}, nil
	}
	// User found. Update Line info
	update := line.LineUser{
		LineServiceID:         lineServiceID,
		LineDisplayName:       profile.DisplayName,
		EncryptedAccessToken:  aToken,
		EncryptedRefreshToken: rToken,
		UserID:                lu.UserID,
	}

	if err := uc.lineC.Set(ipt.Ctx, update); err != nil {
		err = fmt.Errorf("failed to update line user: %v", err)
		log.Printf("error in line auth: %v\n", err)
		return nil, err
	}
	u, err := uc.userQ.GetByID(ipt.Ctx, lu.UserID)
	if err != nil {
		err = fmt.Errorf("authorized, but failed to get user: %v", err)
		log.Printf("error in line auth: %v\n", err)
		return nil, err
	}
	if u.Detail.PictureURL != user.PictureURL(profile.PictureURL) {
		updateDetail := user.Detail{
			PictureURL: user.PictureURL(profile.PictureURL),
		}
		if err := uc.userC.UpdateUserDetail(ipt.Ctx, lu.UserID, updateDetail); err != nil {
			err = fmt.Errorf("failed to update user: %v", err)
			log.Printf("error in line auth: %v\n", err)
			return nil, err
		}
	}
	return &AuthOutput{
		UserInfo:      u,
		LineServiceID: lineServiceID,
	}, nil
}
