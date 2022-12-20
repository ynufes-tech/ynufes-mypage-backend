package line

import (
	"github.com/gin-gonic/gin"
	"github.com/godruoyi/go-snowflake"
	"log"
	"strconv"
	"time"
	"ynufes-mypage-backend/pkg/jwt"
	"ynufes-mypage-backend/pkg/line"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/svc/pkg/config"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type LineAuth struct {
	verifier lineDomain.AuthVerifier
	userQ    query.User
	userC    command.User
	domain   string
}

func NewLineAuth(registry registry.Registry) LineAuth {
	conf := setting.Get()
	return LineAuth{
		verifier: registry.Service().NewLineAuthVerifier(),
		userQ:    registry.Repository().NewUserQuery(),
		userC:    registry.Repository().NewUserCommand(),
		domain:   conf.Application.Server.Domain,
	}
}

func (a LineAuth) VerificationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Request.URL.Query().Get("code")
		state := c.Request.URL.Query().Get("state")
		token, err := a.verifier.RequestAccessToken(code, state)
		if err != nil {
			log.Println("Failed to get access token from LINE server... ", err)
			_, _ = c.Writer.WriteString(err.Error())
			c.AbortWithStatus(500)
			return
		}
		accessToken, err := a.verifier.VerifyAccessToken(token.AccessToken)
		if err != nil {
			log.Println("Failed to verify access token... ", err)
			_, _ = c.Writer.WriteString(err.Error())
			c.AbortWithStatus(500)
			return
		}
		profile, err := line.GetProfile(token.AccessToken)
		if err != nil {
			// failed to get profile
			log.Println(c, "failed to get profile: %v", err)
			c.AbortWithStatus(500)
			return
		}
		u, err := a.userQ.GetByLineServiceID(c, user.LineServiceID(accessToken.ClientId))
		if err != nil {
			// if error is "user not found", Create User and redirect to basic info form
			// Otherwise, respond with error
			newID := snowflake.ID()
			aToken, err := user.NewEncryptedAccessToken(user.PlainAccessToken(token.AccessToken))
			if err != nil {
				log.Println(c, "failed to encrypt access token: %v", err)
				c.AbortWithStatus(500)
				return
			}
			rToken, err := user.NewEncryptedRefreshToken(user.PlainRefreshToken(token.RefreshToken))
			if err != nil {
				log.Println(c, "failed to encrypt refresh token: %v", err)
				c.AbortWithStatus(500)
				return
			}
			err = a.userC.Create(c, user.User{
				ID:     user.ID(newID),
				Status: user.StatusNew,
				Line: user.Line{
					LineServiceID:         user.LineServiceID(accessToken.ClientId),
					LineProfilePictureURL: user.LineProfilePictureURL(profile.PictureURL),
					LineDisplayName:       profile.DisplayName,
					EncryptedAccessToken:  aToken,
					EncryptedRefreshToken: rToken,
				},
			})
			if err != nil {
				c.AbortWithStatus(401)
				return
			}
			err = a.setCookie(c, strconv.FormatUint(newID, 10))
			if err != nil {
				c.AbortWithStatus(500)
				return
			}
			c.Redirect(302, "/welcome")
			return
		}
		// if user exists, update line token, set NewJWT, and redirect to home
		err = a.setCookie(c, strconv.FormatUint(uint64(u.ID), 10))
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		err = a.userC.UpdateLine(c, *u)
		if err != nil {
			log.Println(c, "failed to update line token: %v", err)
			c.AbortWithStatus(500)
			return
		}
		// give JWT and redirect to home
		// if user basic info is not filled, redirect to basic info form
		if u.Status == user.StatusNew {
			c.Redirect(302, "/welcome")
			return
		}
		c.Redirect(302, "/")
	}
}

func (a LineAuth) setCookie(c *gin.Context, id string) error {
	claim := jwt.CreateClaims(id, 24*time.Hour, a.domain)
	token, err := jwt.IssueJWT(claim, config.JWT.JWTSecret)
	if err != nil {
		return err
	}
	// maxAge is set to 1 day
	c.SetCookie("Authorization", "Bearer "+token, 3600*24, "/", a.domain, true, true)
	return nil
}

func (a LineAuth) StateIssuer() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := a.verifier.IssueNewState()
		_, _ = c.Writer.WriteString(state)
		c.Status(200)
	}
}
