package admin

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/registry"
	"ynufes-mypage-backend/svc/pkg/schema/admin"
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
		eventName, has := c.GetQuery("event_name")
		if !has {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "event_name is required",
			})
			return
		}
		opt, err := uc.createUC.Do(event.CreateInput{
			Ctx:       c,
			EventName: eventName,
		})
		if err != nil {
			return
		}
		resp := admin.CreateEventResponse{
			EventID:   opt.Event.ID.ExportID(),
			EventName: opt.Event.Name,
		}
		c.JSON(200, resp)
	}
}
