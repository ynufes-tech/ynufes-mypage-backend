package agent

import (
	"github.com/gin-gonic/gin"
	"strconv"
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
		eventID, err := identity.ImportID(c.Query("event_id"))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid event_id"})
			return
		}
		orgName := c.Query("org_name")
		isOpenRaw := c.DefaultQuery("is_open", "true")
		isOpen, err := strconv.ParseBool(isOpenRaw)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid is_open"})
			return
		}
		ipt := org.CreateOrgInput{
			Ctx:     c,
			EventID: eventID,
			OrgName: orgName,
			IsOpen:  isOpen,
		}
		opt, err := o.createOrgUC.Do(ipt)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to create org"})
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
