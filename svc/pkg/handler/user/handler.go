package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/middleware"
	"ynufes-mypage-backend/svc/pkg/registry"
	uc "ynufes-mypage-backend/svc/pkg/uc/user"
)

type User struct {
	userQ  query.User
	infoUC uc.InfoUseCase
}

func NewUser(rgst registry.Registry) User {
	return User{
		userQ:  rgst.Repository().NewUserQuery(),
		infoUC: uc.NewInfoUseCase(),
	}
}

func (ui User) InfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get(middleware.UserContextKey)
		u = u.(user.User)
		if !exists || !(u).(user.User).IsValid() {
			c.AbortWithError(500, errors.New("failed to retrieve user from context"))
			return
		}
		_, _ = c.Writer.WriteString(ui.infoUC.Do(uc.InfoInput{User: u.(user.User)}).Response)
		c.Status(200)
	}
}
