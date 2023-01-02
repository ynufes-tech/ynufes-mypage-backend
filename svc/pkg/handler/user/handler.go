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
		if !exists || !ok || !u.IsValid() {
			c.AbortWithStatusJSON(400, gin.H{"status": false, "message": "failed to retrieve user from context"})
			return
		}
		var req schema.InfoUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"status": false, "message": err.Error()})
			return
		}
		newDetail := u.Detail
		err := req.ApplyToDetail(&newDetail)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"status": false, "message": err.Error()})
			return
		}
		out, err := uh.infoUpdateUC.Do(uc.UserInfoUpdateInput{
			Ctx:       c,
			OldUser:   &u,
			NewDetail: newDetail,
		})
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"status": false, "message": err.Error()})
			return
		}
		if out.Error != nil {
			c.AbortWithStatusJSON(400, gin.H{"status": false, "message": out.Error.Error()})
			return
		}
		c.Status(200)
	}
}
