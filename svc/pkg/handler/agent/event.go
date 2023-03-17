package agent

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/registry"
	"ynufes-mypage-backend/svc/pkg/schema/agent"
	"ynufes-mypage-backend/svc/pkg/uc/event"
)

type Event struct {
	createUC event.CreateUseCase
}

func NewEvent(rgst registry.Registry) Event {
	return Event{
		createUC: *event.NewCreate(rgst),
	}
}

func (uc Event) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req agent.CreateEventRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "invalid request",
			})
			return
		}
		if req.EventName == "" {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "event_name is required",
			})
			return
		}
		opt, err := uc.createUC.Do(event.CreateInput{
			Ctx:       c,
			EventName: req.EventName,
		})
		if err != nil {
			return
		}
		resp := agent.CreateEventResponse{
			EventID:   opt.Event.ID.ExportID(),
			EventName: opt.Event.Name,
		}
		c.JSON(200, resp)
	}
}
