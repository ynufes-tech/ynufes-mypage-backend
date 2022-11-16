package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	router := gin.Default()
	router.GET("/hello/", testHello)
	router.GET("/auth/line/callback", lineCallback)
	//router.GET("/auth/line/reqState", reqState)
	err := router.Run("localhost:1306")
	if err != nil {
		fmt.Println("Failed to start server...")
	}
}

func loadEnv() {
	godotenv.Load()
}
