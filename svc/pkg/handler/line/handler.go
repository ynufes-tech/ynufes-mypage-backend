package line

import (
	"github.com/gin-gonic/gin"
	"github.com/godruoyi/go-snowflake"
	"log"
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
	verifier   lineDomain.AuthVerifier
	userQ      query.User
	userC      command.User
	domain     string
	devSetting devSetting
}
type devSetting struct {
	callbackURI string
	clientID    string
}

func NewLineAuth(registry registry.Registry) LineAuth {
	conf := setting.Get()
	return LineAuth{
		verifier: registry.Service().NewLineAuthVerifier(),
		userQ:    registry.Repository().NewUserQuery(),
		userC:    registry.Repository().NewUserCommand(),
		domain:   conf.Application.Server.Domain,
		devSetting: devSetting{
			callbackURI: conf.ThirdParty.LineLogin.CallbackURI,
			clientID:    conf.ThirdParty.LineLogin.ClientID,
		},
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
		profile, err := line.GetProfile(token.AccessToken)
		if err != nil {
			// failed to get profile
			log.Println(c, "failed to get profile: %v", err)
			c.AbortWithStatus(500)
			return
		}
		lineServiceID := user.LineServiceID(profile.UserID)
		u, err := a.userQ.GetByLineServiceID(c, lineServiceID)
		if err != nil {
			// if error is "user not found", Create User and redirect to basic info form
			// Otherwise, respond with error
			newID := user.ID(snowflake.ID())
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
				ID:     newID,
				Status: user.StatusNew,
				Line: user.Line{
					LineServiceID:         lineServiceID,
					LineProfilePictureURL: user.LineProfilePictureURL(profile.PictureURL),
					LineDisplayName:       profile.DisplayName,
					EncryptedAccessToken:  aToken,
					EncryptedRefreshToken: rToken,
				},
			})
			if err != nil {
				log.Println(c, "failed to create user: %v", err)
				c.AbortWithStatus(401)
				return
			}
			err = a.setCookie(c, newID.ExportID())
			if err != nil {
				log.Println(c, "failed to set cookie: %v", err)
				c.AbortWithStatus(500)
				return
			}
			c.Redirect(302, "/welcome")
			return
		}
		// if user exists, update line token, set NewJWT, and redirect to home
		err = a.setCookie(c, u.ID.ExportID())
		if err != nil {
			log.Println(c, "failed to set cookie: %v", err)
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
	c.SetCookie("Authorization", token, 3600*24, "/", a.domain, true, true)
	return nil
}

func (a LineAuth) StateIssuer() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := a.verifier.IssueNewState()
		_, _ = c.Writer.WriteString(state)
		c.Status(200)
	}
}

func (a LineAuth) DevAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := a.verifier.IssueNewState()
		c.Redirect(302, "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id="+a.devSetting.clientID+"&redirect_uri="+a.devSetting.callbackURI+"&state="+state+"&scope=openid%20profile%20email")
	}
}
