package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"ynufes-mypage-backend/svc/pkg/handler/test"
	"ynufes-mypage-backend/svc/runner"
)

func main() {
	loadEnv()
	engine := gin.New()
	apiV1 := engine.Group("/api/v1")
	err := runner.Implement(apiV1)
	if err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
	engine.GET("/hello/", test.TestHello)
	err = engine.Run("localhost:1306")
	if err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

//func devAuth(c *gin.Context) {
//	config := setting.Get()
//	c.Redirect(302, "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id="+os.Getenv(config.ThirdParty.LineLogin.ClientID)+
//		"&redirect_uri="+url.QueryEscape(os.Getenv(config.ThirdParty.LineLogin.CallbackURI))+"&state="+linePkg.IssueNewState()+"&scope=openid%20profile%20email")
//}
