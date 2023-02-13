package runner

import (
	"github.com/gin-gonic/gin"
	adminOrg "ynufes-mypage-backend/svc/pkg/handler/admin"
	"ynufes-mypage-backend/svc/pkg/handler/line"
	orgHandler "ynufes-mypage-backend/svc/pkg/handler/org"
	userHandler "ynufes-mypage-backend/svc/pkg/handler/user"
	"ynufes-mypage-backend/svc/pkg/middleware"
	"ynufes-mypage-backend/svc/pkg/registry"
)

func Implement(rg *gin.RouterGroup, devTool bool) error {
	rgst, err := registry.New()
	if err != nil {
		return err
	}
	lineAuth := line.NewLineAuth(*rgst)
	rg.Handle("GET", "/auth/line/callback", lineAuth.VerificationHandler())
	rg.Handle("GET", "/auth/line/state", lineAuth.StateIssuer())

	if devTool {
		//method for development purpose
		rg.Handle("GET", "/auth/line/dev", lineAuth.DevAuth())
		rg.OPTIONS("/*any", func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(200)
		})
	}

	middlewareAuth := middleware.NewAuth(*rgst)

	user := userHandler.NewUser(*rgst)
	org := orgHandler.NewOrg(*rgst)
	authRg := rg.Use(middlewareAuth.VerifyUser())
	authRg.Handle("GET", "/user/info", user.InfoHandler())
	authRg.Handle("POST", "/user/info/update", user.InfoUpdateHandler())
	authRg.Handle("GET", "/orgs", org.OrgsHandler())
	authRg.Handle("POST", "/org/register", org.OrgRegisterHandler())
	return nil
}

func ImplementAdmin(rg *gin.RouterGroup) error {
	// TODO: Implement admin Auth middleware
	rgst, err := registry.New()
	if err != nil {
		return err
	}
	//middlewareAuth := middleware.NewAuth(*rgst)
	//middlewareAdmin := middleware.NewAdmin(*rgst)
	//adminRg := rg.Use(middlewareAuth.VerifyUser(), middlewareAdmin.VerifyAdmin())
	event := adminOrg.NewEvent(*rgst)
	org := adminOrg.NewOrg(*rgst)
	rg.Handle("GET", "/admin/event/create", event.CreateHandler())
	rg.Handle("GET", "/admin/org/create", org.CreateHandler())
	rg.Handle("GET", "/admin/org/token", org.IssueOrgInviteToken())
	return nil
}
