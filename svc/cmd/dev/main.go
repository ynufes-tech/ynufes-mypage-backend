package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"ynufes-mypage-backend/svc/pkg/handler/test"
	"ynufes-mypage-backend/svc/runner"
)

func main() {
	engine := gin.Default()
	apiV1 := engine.Group("/api/v1")
	if err := runner.Implement(apiV1, true); err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
	if err := runner.ImplementAgent(apiV1); err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}

	engine.GET("/hello/", test.TestHello)
	if err := engine.Run("localhost:1306"); err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
}
