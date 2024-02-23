package token

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/domain/service/auth"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type Token struct {
	issuer *auth.TokenIssuer
}

func NewToken(r registry.Registry) Token {
	return Token{
		issuer: r.Service().TokenIssuer(),
	}
}

// IssueHandler works as post handler
func (t Token) IssueHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.PostForm("code")
		if code == "" {
			c.JSON(400, gin.H{"error": "code is required"})
			return
		}
		token, err := (*t.issuer).IssueToken(code)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"token": token})
	}
}
