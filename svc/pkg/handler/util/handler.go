package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/middleware"
)

type Handler func(c *gin.Context, user user.User)

func (h Handler) GinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		uAny, exists := c.Get(middleware.UserContextKey)
		if !exists || !(uAny).(user.User).IsValid() {
			_ = c.AbortWithError(500, errors.New("failed to retrieve user from context"))
			return
		}
		h(c, uAny.(user.User))
	}
}
