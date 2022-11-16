package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ynufes-mypage-backend/pkg/line"
)

func testHello(c *gin.Context) {
	name := c.Request.URL.Path[len("/hello/"):]
	c.Writer.WriteString("Hello " + name)
}

func lineCallback(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	state := c.Request.URL.Query().Get("state")
	if !line.VerifyState(state) {
		c.Status(401)
		c.Writer.WriteString("Your request is not valid...")
		return
	}
	accessResponse, err := line.RequestAccessToken(code)
	if err != nil {
		c.Writer.WriteString("Error, " + err.Error())
	} else {
		c.Writer.WriteString("AccessToken: " + accessResponse.AccessToken + "\n")
		c.Writer.WriteString("TokenType: " + accessResponse.TokenType + "\n")
		c.Writer.WriteString("ExpiresIn: " + strconv.FormatInt(accessResponse.ExpiresIn, 10) + "\n")
		c.Writer.WriteString("RefreshToken: " + accessResponse.RefreshToken + "\n")
		c.Writer.WriteString("Scope: " + accessResponse.Scope + "\n")
	}

}
