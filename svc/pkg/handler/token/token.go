package token

import (
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/svc/pkg/domain/service/auth"
	"ynufes-mypage-backend/svc/pkg/registry"
	schema "ynufes-mypage-backend/svc/pkg/schema/token"
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
		var req schema.TokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Code == "" {
			c.JSON(400, gin.H{"error": "code is required"})
			return
		}
		token, err := (*t.issuer).IssueToken(req.Code)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, schema.TokenResponse{token})
	}
}
