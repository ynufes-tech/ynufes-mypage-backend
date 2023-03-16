package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/handler/util"
	"ynufes-mypage-backend/svc/pkg/registry"
	userSchema "ynufes-mypage-backend/svc/pkg/schema/user"
	uc "ynufes-mypage-backend/svc/pkg/uc/user"
)

type User struct {
	infoUpdateUC uc.UserInfoUpdateUseCase
}

func NewUser(rgst registry.Registry) User {
	return User{
		infoUpdateUC: uc.NewInfoUpdate(rgst),
	}
}

func (uh User) InfoHandler() gin.HandlerFunc {
	var h util.Handler = func(c *gin.Context, u user.User) {
		var status int
		if u.Detail.MeetsBasicRequirement() {
			status = 2
		} else {
			status = 1
		}
		output := userSchema.InfoResponse{
			NameFirst:       u.Detail.Name.FirstName,
			NameLast:        u.Detail.Name.LastName,
			NameFirstKana:   u.Detail.Name.FirstNameKana,
			NameLastKana:    u.Detail.Name.LastNameKana,
			Type:            int(u.Detail.Type),
			ProfileImageURL: string(u.Detail.PictureURL),
			Email:           string(u.Detail.Email),
			Gender:          int(u.Detail.Gender),
			StudentID:       string(u.Detail.StudentID),
			Status:          status,
		}
		j, err := json.Marshal(output)
		if err != nil {
			_ = c.AbortWithError(500, fmt.Errorf("failed to marshal output: %w", err))
			return
		}
		_, _ = c.Writer.WriteString(string(j))
		c.Status(200)
	}
	return h.GinHandler()
}

func (uh User) InfoUpdateHandler() gin.HandlerFunc {
	var h util.Handler = func(c *gin.Context, u user.User) {
		var req userSchema.InfoUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"status": false, "message": err.Error()})
			return
		}
		err := req.ApplyToDetail(&u.Detail)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"status": false, "message": err.Error()})
			return
		}
		out, err := uh.infoUpdateUC.Do(uc.UserInfoUpdateInput{
			Ctx:     c,
			NewUser: u,
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
	return h.GinHandler()
}
