package runner

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/handler/line"
	"ynufes-mypage-backend/svc/pkg/registry"
)

func Implement(rg *gin.RouterGroup) error {
	rgst, err := registry.New()
	if err != nil {
		return err
	}
	lineAuth := line.NewLineAuth(*rgst)
	rg.Handle("GET", "/auth/line/verify", lineAuth.VerificationHandler())
	rg.Handle("GET", "/auth/line/state", lineAuth.StateIssuer())
	return nil
}
