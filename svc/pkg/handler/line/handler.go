package line

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/domain/service/auth"
	lineDomain "ynufes-mypage-backend/svc/pkg/domain/service/line"
	"ynufes-mypage-backend/svc/pkg/registry"
	lineUC "ynufes-mypage-backend/svc/pkg/uc/line"
)

type LineAuth struct {
	verifier     *lineDomain.AuthVerifier
	userQ        query.User
	userC        command.User
	serverConf   setting.Server
	devSetting   devSetting
	secureCookie bool
	authUC       lineUC.AuthUseCase
	tokenIssuer  *auth.TokenIssuer
}

type devSetting struct {
	callbackURI string
	clientID    string
}

func NewLineAuth(registry registry.Registry) LineAuth {
	conf := setting.Get()
	authVerifier := registry.Service().LineAuthVerifier()
	return LineAuth{
		verifier:   authVerifier,
		userQ:      registry.Repository().NewUserQuery(),
		userC:      registry.Repository().NewUserCommand(),
		serverConf: conf.Application.Server,
		devSetting: devSetting{
			callbackURI: conf.ThirdParty.LineLogin.CallbackURI,
			clientID:    conf.ThirdParty.LineLogin.ClientID,
		},
		secureCookie: conf.Service.Authentication.SecureCookie,
		authUC:       lineUC.NewAuthCodeUseCase(registry, conf.ThirdParty.LineLogin.EnableLineAuth, authVerifier),
		tokenIssuer:  registry.Service().TokenIssuer(),
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

		var redirectDest string
		if a.serverConf.OnProduction {
			redirectDest = "/token"
		} else {
			redirectDest = fmt.Sprintf(
				"%s%s%s/token", a.serverConf.Frontend.Protocol, a.serverConf.Frontend.Domain, a.serverConf.Frontend.Port,
			)
		}
		redirectDest, err = a.attachCode(authOut.UserInfo.ID.ExportID(), redirectDest)
		if err != nil {
			log.Printf("failed to attach token: %v", err)
			c.AbortWithStatus(500)
			return
		}
		c.Redirect(302, redirectDest)
	}
}

func (a LineAuth) attachCode(id string, dest string) (string, error) {
	code, err := (*a.tokenIssuer).IssueNewCode(id)
	if err != nil {
		return "", fmt.Errorf("failed to issue code: %v", err)
	}
	dest = fmt.Sprintf("%s?code=%s", dest, code)
	return dest, nil
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
