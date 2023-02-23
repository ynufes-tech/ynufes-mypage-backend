package agent

import (
	"github.com/gin-gonic/gin"
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/handler/util"
	"ynufes-mypage-backend/svc/pkg/registry"
	"ynufes-mypage-backend/svc/pkg/schema/agent"
	"ynufes-mypage-backend/svc/pkg/uc/form"
)

type Form struct {
	createUC form.CreateUseCase
}

func NewForm(rgst registry.Registry) Form {
	return Form{
		createUC: *form.NewCreate(rgst),
	}
}

func (h Form) CreateHandler() gin.HandlerFunc {
	uh := util.Handler(func(c *gin.Context, targetUser user.User) {
		var req agent.CreateFormRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "invalid request body",
			})
			return
		}
		eventID, err := identity.ImportID(req.EventID)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "invalid event_id",
			})
			return
		}
		deadline := time.UnixMilli(req.Deadline)
		if deadline.Before(time.Now()) {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "invalid deadline",
			})
			return
		}
		if len(req.Title) < 5 {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "title is too short",
			})
			return
		}
		ipt := form.CreateInput{
			Ctx:         c,
			User:        targetUser,
			EventID:     eventID,
			Title:       req.Title,
			Summary:     req.Summary,
			Description: req.Description,
			Deadline:    deadline,
		}
		opt, err := h.createUC.Do(ipt)
		if err != nil {
			c.AbortWithStatusJSON(500,
				gin.H{"error": "failed to create form"})
			_ = c.Error(err)
			return
		}
		c.JSON(200, gin.H{
			"form_id": opt.FormID.ExportID(),
		})
	})
	return uh.GinHandler()
}
