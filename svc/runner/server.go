package runner

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/handler/line"
	"ynufes-mypage-backend/svc/pkg/registry"
)

func New() (*gin.Engine, error) {
	rgst, err := registry.New()
	if err != nil {
		return nil, err
	}
	engine := gin.New()
	lineAuth := line.NewLineAuth(*rgst)
	engine.Handle("GET", "/auth/line/verify", lineAuth.VerificationHandler())
	engine.Handle("GET", "/auth/line/state", lineAuth.StateIssuer())
	return engine, nil
}
