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
	orgsUC     org.OrgsUseCase
	registerUC org.RegisterUseCase
}

func NewOrg(rgst registry.Registry) Org {
	return Org{
		orgsUC:     org.NewOrgs(rgst),
		registerUC: org.NewRegister(rgst),
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

func (o Org) OrgRegisterHandler() gin.HandlerFunc {
	var h util.Handler = func(ctx *gin.Context, user user.User) {
		var req schema.RegisterRequest
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid request"})
			return
		}
		ipt := org.RegisterInput{
			Ctx:    ctx,
			UserID: user.ID,
			Token:  req.Token,
		}
		opt, err := o.registerUC.Do(ipt)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		resp := schema.RegisterResponse{
			Added:     opt.Added,
			OrgID:     opt.Org.ID.ExportID(),
			OrgName:   opt.Org.Name,
			EventID:   opt.Org.Event.ID.ExportID(),
			EventName: opt.Org.Event.Name,
		}
		ctx.JSON(200, resp)
	}
	return h.GinHandler()
}
