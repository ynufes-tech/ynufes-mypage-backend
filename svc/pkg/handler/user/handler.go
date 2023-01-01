package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/middleware"
	"ynufes-mypage-backend/svc/pkg/registry"
	"ynufes-mypage-backend/svc/pkg/schema"
	uc "ynufes-mypage-backend/svc/pkg/uc/user"
)

type User struct {
	userQ        query.User
	userC        command.User
	infoUC       uc.InfoUseCase
	infoUpdateUC uc.UserInfoUpdateUseCase
}

func NewUser(rgst registry.Registry) User {
	return User{
		userQ:        rgst.Repository().NewUserQuery(),
		userC:        rgst.Repository().NewUserCommand(),
		infoUC:       uc.NewInfoUseCase(),
		infoUpdateUC: uc.NewInfoUpdate(rgst),
	}
}

func (uh User) InfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get(middleware.UserContextKey)
		u = u.(user.User)
		if !exists || !(u).(user.User).IsValid() {
			_ = c.AbortWithError(500, errors.New("failed to retrieve user from context"))
			return
		}
		_, _ = c.Writer.WriteString(uh.infoUC.Do(uc.InfoInput{User: u.(user.User)}).Response)
		c.Status(200)
	}
}

func (uh User) InfoUpdateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		uA, exists := c.Get(middleware.UserContextKey)
		u, ok := uA.(user.User)
		if !exists || !ok || u.IsValid() {
			_ = c.AbortWithError(500, errors.New("failed to retrieve user from context"))
			return
		}
		var req schema.InfoUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			_ = c.AbortWithError(400, err)
			return
		}
		detail, err := req.ToUserDetail()
		if err != nil {
			_ = c.AbortWithError(400, err)
			return
		}
		out := uh.infoUpdateUC.Do(uc.UserInfoUpdateInput{
			OldUser:   &u,
			NewDetail: *detail,
		})
		if out.Error != nil {
			_ = c.AbortWithError(500, out.Error)
		}
		c.Status(200)
	}
}
