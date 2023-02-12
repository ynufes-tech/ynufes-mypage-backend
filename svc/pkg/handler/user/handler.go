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
		output := userSchema.InfoResponse{
			NameFirst:       u.Detail.Name.FirstName,
			NameLast:        u.Detail.Name.LastName,
			Type:            int(u.Detail.Type),
			ProfileImageURL: string(u.Line.LineProfilePictureURL),
			Status:          int(u.Status),
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
	return h.GinHandler()
}
