package runner

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/handler/line"
	userHandler "ynufes-mypage-backend/svc/pkg/handler/user"
	"ynufes-mypage-backend/svc/pkg/middleware"
	"ynufes-mypage-backend/svc/pkg/registry"
)

func Implement(rg *gin.RouterGroup) error {
	rgst, err := registry.New()
	if err != nil {
		return err
	}
	lineAuth := line.NewLineAuth(*rgst)
	rg.Handle("GET", "/auth/line/callback", lineAuth.VerificationHandler())
	rg.Handle("GET", "/auth/line/state", lineAuth.StateIssuer())

	//method for development purpose
	rg.Handle("GET", "/auth/line/dev", lineAuth.DevAuth())

	middlewareAuth := middleware.NewAuth(*rgst)

	user := userHandler.NewUser(*rgst)
	authRg := rg.Use(middlewareAuth.VerifyUser())
	authRg.Handle("GET", "/user/info", user.InfoHandler())
	return nil
}
