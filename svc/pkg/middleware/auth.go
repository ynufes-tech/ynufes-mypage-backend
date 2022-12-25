package middleware

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/registry"
	uc "ynufes-mypage-backend/svc/pkg/uc/user"
)

const (
	AuthorizedUserIDField = "AuthorizedUserID"
	UserContextKey        = "User"
)

type auth struct {
	loginUC uc.LoginUseCase
}

func NewAuth(registry registry.Registry) auth {
	return auth{
		loginUC: uc.NewLoginUseCase(registry),
	}
}

func (a auth) VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := getJWTFromHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		result, err := a.loginUC.Do(c, uc.LoginInput{JWT: user.JWT(jwt)})
		c.Set(UserContextKey, result)
		if result == nil || err != nil {
			c.AbortWithError(401, err)
		}
		c.Set(AuthorizedUserIDField, result.User.ID)
		c.Next()
	}
}

func getJWTFromHeader(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")

	if len(header) < 8 || header[:7] != "Bearer " {
		return "", exception.ErrorInvalidHeader
	}
	return header[7:], nil
}
