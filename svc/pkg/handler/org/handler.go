package org

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/handler/util"
	"ynufes-mypage-backend/svc/pkg/registry"
	schemaForm "ynufes-mypage-backend/svc/pkg/schema/form"
	schemaOrg "ynufes-mypage-backend/svc/pkg/schema/org"
	"ynufes-mypage-backend/svc/pkg/uc/org"
)

type Org struct {
	orgsUC     org.OrgsUseCase
	registerUC org.RegisterUseCase
	orgUC      org.OrgUseCase
}

func NewOrg(rgst registry.Registry) Org {
	return Org{
		orgsUC:     org.NewOrgs(rgst),
		registerUC: org.NewRegister(rgst),
		orgUC:      org.NewOrg(rgst),
	}
}

// OrgsHandler returns a list of organizations that the user is registered to.
func (o Org) OrgsHandler() gin.HandlerFunc {
	var h util.Handler = func(context *gin.Context, u user.User) {
		ipt := org.OrgsInput{
			Ctx:    context,
			UserID: u.ID,
		}
		opt, err := o.orgsUC.Do(ipt)
		if err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			log.Printf("error in OrgHandler: %v", err)
			return
		}
		orgs := make([]schemaOrg.Org, len(opt.Orgs))
		for i, or := range opt.Orgs {
			orgs[i] = schemaOrg.Org{
				ID:        or.ID.ExportID(),
				Name:      or.Name,
				EventName: or.Event.Name,
				EventID:   or.Event.ID.ExportID(),
				IsOpen:    or.IsOpen,
			}
		}
		resp := schemaOrg.OrgsResponse{
			Orgs: orgs,
		}
		context.JSON(200, resp)
	}
	return h.GinHandler()
}

// OrgRegisterHandler accepts a request to register a user to an organization.
// The request must contain a token, which is generated by appropriate agents.
// Tokens can be issued by appropriate agents with /agent/org/token endpoint.
func (o Org) OrgRegisterHandler() gin.HandlerFunc {
	var h util.Handler = func(ctx *gin.Context, u user.User) {
		var req schemaOrg.RegisterRequest
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid request"})
			return
		}
		ipt := org.RegisterInput{
			Ctx:    ctx,
			UserID: u.ID,
			Token:  req.Token,
		}
		opt, err := o.registerUC.Do(ipt)
		if err != nil {
			if errors.Is(err, exception.ErrNotFound) {
				ctx.AbortWithStatusJSON(404, gin.H{"error": "org not found"})
				return
			}
			log.Printf("error in OrgRegisterHandler: %v\n", err)
			_ = ctx.AbortWithError(500, err)
			return
		}
		resp := schemaOrg.RegisterResponse{
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

// OrgHandler returns information required for Org page.
// /org/:orgID
func (o Org) OrgHandler() gin.HandlerFunc {
	var h util.Handler = func(ctx *gin.Context, u user.User) {
		orgID, err := identity.ImportID(ctx.Param("orgID"))
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid orgID"})
			return
		}
		ipt := org.OrgInput{
			Ctx:   ctx,
			User:  u,
			OrgID: orgID,
		}
		opt, err := o.orgUC.Do(ipt)
		if err != nil {
			if errors.Is(err, exception.ErrNotFound) {
				ctx.AbortWithStatusJSON(404, gin.H{"error": "org not found"})
				return
			}
			log.Printf("error in OrgHandler: %v\n", err)
			_ = ctx.AbortWithError(500, err)
			return
		}
		forms := make([]schemaForm.FormSummary, len(opt.Forms))
		for i, f := range opt.Forms {
			forms[i] = schemaForm.FormSummary{
				ID:          f.ID.ExportID(),
				Title:       f.Title,
				Summary:     f.Summary,
				Description: f.Description,
				Deadline:    f.Deadline.Format(time.RFC3339),
				Status:      int(form.Accepted),
				IsOpen:      f.IsOpen,
			}
		}
		resp := schemaOrg.NewOrgResponse(
			opt.Org.ID.ExportID(),
			opt.Org.Name,
			opt.Org.Event.ID.ExportID(),
			opt.Org.Event.Name,
			forms,
		)
		ctx.JSON(200, resp)
	}
	return h.GinHandler()
}
