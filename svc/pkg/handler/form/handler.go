package form

import (
	"github.com/gin-gonic/gin"
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/handler/util"
	"ynufes-mypage-backend/svc/pkg/registry"
	formResp "ynufes-mypage-backend/svc/pkg/schema/form"
	"ynufes-mypage-backend/svc/pkg/uc/form"
)

type Form struct {
	infoUC form.InfoUseCase
}

func NewForm(rgst registry.Registry) Form {
	return Form{
		infoUC: form.NewInfo(rgst),
	}
}

func (f Form) InfoHandler() gin.HandlerFunc {
	return util.Handler(func(ctx *gin.Context, u user.User) {
		orgSID, has := ctx.GetQuery("org_id")
		if !has {
			ctx.JSON(400, gin.H{"error": "org_id is required"})
			return
		}
		formSID, has := ctx.GetQuery("form_id")
		if !has {
			ctx.JSON(400, gin.H{"error": "form_id is required"})
			return
		}
		orgID, err := identity.ImportID(orgSID)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		formID, err := identity.ImportID(formSID)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ipt := form.InfoInput{
			Ctx:    ctx,
			UserID: u.ID,
			OrgID:  orgID,
			FormID: formID,
		}
		opt, err := f.infoUC.Do(ipt)
		if err != nil {
			if err == exception.ErrUnauthorized {
				ctx.JSON(401, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		secsO := opt.Form.Sections.GetOrderedIDs()
		secs := make([]string, len(opt.Form.Sections))
		for i := range secsO {
			secs[i] = secsO[i].ExportID()
		}
		resp := formResp.FormInfoResponse{
			ID:          opt.Form.ID.ExportID(),
			Title:       opt.Form.Title,
			Summary:     opt.Form.Summary,
			Description: opt.Form.Description,
			Deadline:    opt.Form.Deadline.Format(time.RFC3339Nano),
			// TODO: Set status
			Status:   0,
			IsOpen:   opt.Form.IsOpen,
			Sections: secs,
		}
		ctx.JSON(200, resp)
	}).GinHandler()
}
