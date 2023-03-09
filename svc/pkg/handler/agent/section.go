package agent

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/handler/util"
	"ynufes-mypage-backend/svc/pkg/registry"
	"ynufes-mypage-backend/svc/pkg/schema/section"
	uc "ynufes-mypage-backend/svc/pkg/uc/section"
)

type Section struct {
	createUC uc.CreateUseCase
}

func NewSection(rgst registry.Registry) Section {
	return Section{
		createUC: uc.NewCreate(rgst),
	}
}

func (h Section) CreateHandler() gin.HandlerFunc {
	handler := util.Handler(func(c *gin.Context, user user.User) {
		var req section.CreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		fid, err := identity.ImportID(req.FormID)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ipt := uc.CreateInput{
			Ctx:    c,
			UserID: user.ID,
			FormID: fid,
		}
		opt, err := h.createUC.Do(ipt)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(200, section.CreateResponse{
			SectionID: opt.SectionID.ExportID(),
		})
		return
	})
	return handler.GinHandler()
}
