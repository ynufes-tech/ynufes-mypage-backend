package line

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"ynufes-mypage-backend/pkg/jwt"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/svc/pkg/config"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
	"ynufes-mypage-backend/svc/pkg/registry"
	lineUC "ynufes-mypage-backend/svc/pkg/uc/line"
)

type LineAuth struct {
	verifier     *lineDomain.AuthVerifier
	userQ        query.User
	userC        command.User
	domain       string
	devSetting   devSetting
	secureCookie bool
	authUC       lineUC.AuthUseCase
}
type devSetting struct {
	callbackURI string
	clientID    string
}

func NewLineAuth(registry registry.Registry) LineAuth {
	conf := setting.Get()
	authVerifier := registry.Service().NewLineAuthVerifier()
	return LineAuth{
		verifier: &authVerifier,
		userQ:    registry.Repository().NewUserQuery(),
		userC:    registry.Repository().NewUserCommand(),
		domain:   conf.Application.Server.Domain,
		devSetting: devSetting{
			callbackURI: conf.ThirdParty.LineLogin.CallbackURI,
			clientID:    conf.ThirdParty.LineLogin.ClientID,
		},
		secureCookie: conf.Service.Authentication.SecureCookie,
		authUC:       lineUC.NewAuthCodeUseCase(registry, conf.ThirdParty.LineLogin.EnableLineAuth, &authVerifier),
	}
}

func (a LineAuth) VerificationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Request.URL.Query().Get("code")
		state := c.Request.URL.Query().Get("state")
		authInput := lineUC.AuthInput{
			State: state,
			Code:  code,
			Ctx:   c,
		}
		authOut, err := a.authUC.Do(authInput)

		if err != nil {
			_, _ = c.Writer.WriteString(err.Error())
			log.Printf("error: %v", err)
			c.AbortWithStatus(500)
			return
		}
		if authOut.UserInfo == nil {
			_, _ = c.Writer.WriteString(authOut.ErrorMsg)
			c.AbortWithStatus(400)
			return
		}
		err = a.setCookie(c, authOut.UserInfo.ID.ExportID())
		if err != nil {
			log.Println(c, "failed to set cookie: %v", err)
			c.AbortWithStatus(500)
			return
		}
		if authOut.UserInfo.Status == user.StatusNew {
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
	c.SetCookie("Authorization", token, 3600*24, "/", a.domain, a.secureCookie, true)
	return nil
}

func (a LineAuth) StateIssuer() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := (*a.verifier).IssueNewState()
		_, _ = c.Writer.WriteString(state)
		c.Status(200)
	}
}

func (a LineAuth) DevAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := (*a.verifier).IssueNewState()
		c.Redirect(302, "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id="+a.devSetting.clientID+"&redirect_uri="+a.devSetting.callbackURI+"&state="+state+"&scope=openid%20profile%20email")
	}
}
