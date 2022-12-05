package line

import (
	"github.com/gin-gonic/gin"
	"strconv"
	linePkg "ynufes-mypage-backend/pkg/line"
)

func Callback(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	state := c.Request.URL.Query().Get("state")
	if !linePkg.VerifyState(state) {
		c.Status(401)
		c.Writer.WriteString("Your request is not valid...")
		return
	}
	accessResponse, err := linePkg.RequestAccessToken(code)
	if err != nil {
		c.Writer.WriteString("Error, " + err.Error())
	} else {
		c.Writer.WriteString("AccessToken: " + accessResponse.AccessToken + "\n")
		c.Writer.WriteString("TokenType: " + accessResponse.TokenType + "\n")
		c.Writer.WriteString("ExpiresIn: " + strconv.FormatInt(accessResponse.ExpiresIn, 10) + "\n")
		c.Writer.WriteString("RefreshToken: " + accessResponse.RefreshToken + "\n")
		c.Writer.WriteString("Scope: " + accessResponse.Scope + "\n")
	}
	c.SetCookie("access_token", accessResponse.AccessToken, 3600, "/", "ynufes-mypage.shion.pro", true, true)
}
