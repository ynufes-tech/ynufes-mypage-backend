package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/url"
	"os"
	linePkg "ynufes-mypage-backend/pkg/line"
	lineHandler "ynufes-mypage-backend/svc/handler/line"
	"ynufes-mypage-backend/svc/handler/test"
)

func main() {
	loadEnv()
	router := gin.Default()
	router.GET("/hello/", test.TestHello)
	router.GET("/auth/lineHandler/callback", lineHandler.Callback)
	router.GET("/auth/lineHandler/reqState", linePkg.ReqState)
	router.GET("/auth/lineHandler/dev/auth", devAuth)
	err := router.Run("localhost:1306")
	if err != nil {
		fmt.Println("Failed to start server...")
	}
}

func loadEnv() {
	godotenv.Load()
}

func devAuth(c *gin.Context) {
	c.Redirect(302, "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id="+os.Getenv(linePkg.EnvLineClientId)+
		"&redirect_uri="+url.QueryEscape(os.Getenv(linePkg.EnvLineRedirectUri))+"&state="+linePkg.IssueNewState()+"&scope=openid%20profile%20email")
}
