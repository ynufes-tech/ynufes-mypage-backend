package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"ynufes-mypage-backend/pkg/line"
)

func main() {
	loadEnv()
	router := gin.Default()
	router.GET("/hello/", testHello)
	router.GET("/auth/line/callback", lineCallback)
	router.GET("/auth/line/reqState", line.ReqState)
	router.GET("auth/line/dev/auth", devAuth)
	err := router.Run("localhost:1306")
	if err != nil {
		fmt.Println("Failed to start server...")
	}
}

func loadEnv() {
	godotenv.Load()
}

func devAuth(c *gin.Context) {
	c.Redirect(502, "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id="+os.Getenv(line.EnvLineClientId)+
		"&redirect_uri="+os.Getenv(line.EnvLineRedirectUri)+"&state="+line.IssueNewState()+"&scope=openid%20profile%20email")
}
