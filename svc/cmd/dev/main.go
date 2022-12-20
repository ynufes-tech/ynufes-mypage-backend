package main

import (
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/appengine/log"
	"ynufes-mypage-backend/svc/pkg/handler/test"
	"ynufes-mypage-backend/svc/runner"
)

func main() {
	loadEnv()
	engine, err := runner.New()
	if err != nil {
		log.Errorf(context.Background(), "Failed to start server...")
		return
	}
	engine.GET("/hello/", test.TestHello)
	err = engine.Run("localhost:1306")
	if err != nil {
		log.Errorf(context.Background(), "Failed to start server...")
	}
}

func loadEnv() {
	godotenv.Load()
}

//func devAuth(c *gin.Context) {
//	config := setting.Get()
//	c.Redirect(302, "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id="+os.Getenv(config.ThirdParty.LineLogin.ClientID)+
//		"&redirect_uri="+url.QueryEscape(os.Getenv(config.ThirdParty.LineLogin.CallbackURI))+"&state="+linePkg.IssueNewState()+"&scope=openid%20profile%20email")
//}
