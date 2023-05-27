package agent

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/jwt"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/svc/pkg/registry"
	"ynufes-mypage-backend/svc/pkg/schema/agent"
	"ynufes-mypage-backend/svc/pkg/uc/org"
)

type Org struct {
	createOrgUC org.CreateOrgUseCase
	infoOrgUC   org.InfoUseCase
	jwtSecret   string
}

func NewOrg(rgst registry.Registry) *Org {
	config := setting.Get()
	return &Org{
		createOrgUC: org.NewCreateOrg(rgst),
		infoOrgUC:   org.NewInfo(rgst),
		jwtSecret:   config.Application.Authentication.JwtSecret,
	}
}

func (o Org) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req agent.CreateOrgRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid request"})
			return
		}
		if req.OrgName == "" || req.EventID == "" {
			c.AbortWithStatusJSON(400,
				gin.H{"error": "org_name and event_id are required"})
			return
		}
		eventID, err := identity.ImportID(req.EventID)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid event_id"})
			return
		}
		ipt := org.CreateOrgInput{
			Ctx:     c,
			EventID: eventID,
			OrgName: req.OrgName,
			IsOpen:  req.IsOpen,
		}
		opt, err := o.createOrgUC.Do(ipt)
		if err != nil {
			log.Printf("failed to create org in CreateOrgHandler: %v", err)
			c.AbortWithStatusJSON(500,
				gin.H{"error": fmt.Sprintf("failed to create org: %v", err)})
			return
		}
		resp := agent.CreateOrgResponse{
			EventID:   opt.Org.Event.ID.ExportID(),
			EventName: opt.Org.Event.Name,
			OrgID:     opt.Org.ID.ExportID(),
			OrgName:   opt.Org.Name,
			IsOpen:    opt.Org.IsOpen,
		}
		c.JSON(200, resp)
	}
}

// IssueOrgInviteToken issues a JWT token for inviting a user to an org.
func (o Org) IssueOrgInviteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		orgID, err := identity.ImportID(c.Query("org_id"))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid event_id"})
			return
		}
		opt, err := o.infoOrgUC.Do(org.InfoInput{
			Ctx: c,
			ID:  orgID,
		})
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to get org info"})
			return
		}
		durationRaw := c.DefaultQuery("duration", "1h")
		duration, err := time.ParseDuration(durationRaw)
		if err != nil || duration < 0 || duration > 24*time.Hour {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid duration"})
			return
		}
		// TODO: include information about issuer
		claims := jwt.CreateClaims(orgID.ExportID(), duration, "YNUFesMyPageSystem")
		issueJWT, err := jwt.IssueJWT(claims, o.jwtSecret)
		if err != nil {
			return
		}
		resp := agent.IssueOrgInviteTokenResponse{
			Token:      issueJWT,
			OrgID:      opt.Org.ID.ExportID(),
			ValidUntil: time.Now().Add(duration).Format(time.RFC3339),
		}
		c.JSON(200, resp)
	}
}
