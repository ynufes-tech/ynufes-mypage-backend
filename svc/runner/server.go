package runner

import (
	"github.com/gin-gonic/gin"
	agentOrg "ynufes-mypage-backend/svc/pkg/handler/agent"
	formHandler "ynufes-mypage-backend/svc/pkg/handler/form"
	"ynufes-mypage-backend/svc/pkg/handler/line"
	orgHandler "ynufes-mypage-backend/svc/pkg/handler/org"
	sectionHandler "ynufes-mypage-backend/svc/pkg/handler/section"
	"ynufes-mypage-backend/svc/pkg/handler/token"
	userHandler "ynufes-mypage-backend/svc/pkg/handler/user"
	"ynufes-mypage-backend/svc/pkg/middleware"
	"ynufes-mypage-backend/svc/pkg/registry"
)

func Implement(rg *gin.RouterGroup, devTool bool) error {
	rgst, err := registry.New()
	if err != nil {
		return err
	}

	tkn := token.NewToken(*rgst)
	rg.Handle("POST", "/auth/token", tkn.IssueHandler())

	lineAuth := line.NewLineAuth(*rgst)
	rg.Handle("GET", "/auth/line/callback", lineAuth.VerificationHandler())
	rg.Handle("GET", "/auth/line/state", lineAuth.StateIssuer())

	if devTool {
		//method for development purpose
		rg.Handle("GET", "/auth/line/dev", lineAuth.DevAuth())
	}

	middlewareAuth := middleware.NewAuth(*rgst)

	userH := userHandler.NewUser(*rgst)
	orgH := orgHandler.NewOrg(*rgst)
	secH := sectionHandler.NewSection(*rgst)
	formH := formHandler.NewForm(*rgst)
	authRg := rg.Use(middlewareAuth.VerifyUser())
	authRg.Handle("GET", "/user/info", userH.InfoHandler())
	authRg.Handle("POST", "/user/info", userH.InfoUpdateHandler())
	authRg.Handle("GET", "/orgs", orgH.OrgsHandler())
	authRg.Handle("POST", "/org/register", orgH.OrgRegisterHandler())
	authRg.Handle("GET", "/org", orgH.OrgHandler())
	authRg.Handle("GET", "/form", formH.InfoHandler())
	authRg.Handle("GET", "/section", secH.InfoHandler())
	return nil
}

func ImplementAgent(rg *gin.RouterGroup) error {
	// TODO: Implement agent Auth middleware
	rgst, err := registry.New()
	if err != nil {
		return err
	}
	//middlewareAuth := middleware.NewAuth(*rgst)
	//middlewareAgent := middleware.NewAgent(*rgst)
	//agentRg := rg.Use(middlewareAuth.VerifyUser(), middlewareAgent.VerifyAgent())
	event := agentOrg.NewEvent(*rgst)
	org := agentOrg.NewOrg(*rgst)
	form := agentOrg.NewForm(*rgst)
	section := agentOrg.NewSection(*rgst)
	question := agentOrg.NewQuestion(*rgst)
	rg.Handle("POST", "/agent/event/create", event.CreateHandler())
	rg.Handle("POST", "/agent/org/create", org.CreateHandler())
	rg.Handle("GET", "/agent/org/token", org.IssueOrgInviteToken())
	rg.Handle("POST", "/agent/form/create", form.CreateHandler())
	rg.Handle("POST", "/agent/section/create", section.CreateHandler())
	rg.Handle("POST", "/agent/question/create", question.CreateHandler())
	return nil
}
