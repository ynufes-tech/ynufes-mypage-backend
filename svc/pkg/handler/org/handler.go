package org

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/handler/util"
	"ynufes-mypage-backend/svc/pkg/registry"
	schema "ynufes-mypage-backend/svc/pkg/schema/org"
	"ynufes-mypage-backend/svc/pkg/uc/org"
)

type Org struct {
	orgsUC org.OrgsUseCase
}

func NewOrg(rgst registry.Registry) Org {
	return Org{
		orgsUC: org.NewOrgs(rgst),
	}
}

func (o Org) OrgsHandler() gin.HandlerFunc {
	var h util.Handler = func(context *gin.Context, user user.User) {
		ipt := org.OrgsInput{
			Ctx:    context,
			UserID: user.ID,
		}
		opt, err := o.orgsUC.Do(ipt)
		if err != nil {
			context.JSON(500, err)
			return
		}
		orgs := make([]schema.Org, len(opt.Orgs))
		for i, or := range opt.Orgs {
			orgs[i] = schema.Org{
				ID:        or.ID.ExportID(),
				Name:      or.Name,
				EventName: or.Event.Name,
				EventID:   or.Event.ID.ExportID(),
				IsOpen:    or.IsOpen,
			}
		}
		resp := schema.OrgsResponse{
			Orgs: orgs,
		}
		context.JSON(200, resp)
	}
	return h.GinHandler()
}
