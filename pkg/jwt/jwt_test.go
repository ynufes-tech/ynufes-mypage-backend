package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/exception"
)

func TestJWT(t *testing.T) {
	t.Run("JWT_INVALID", func(t *testing.T) {
		claims := jwt.StandardClaims{
			Audience:  "TestAudience",
			ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(),
			Id:        "testID1234",
			IssuedAt:  time.Now().Add(-5 * time.Hour).Unix(),
			Issuer:    "TestIssuer",
			Subject:   "TestSubject",
		}
		issueJWT, err := IssueJWT(claims, "testSecret")
		assert.NoError(t, err)
		newClaims, err := Verify(user.JWT(issueJWT), "testSecret")
		assert.EqualError(t, err, exception.ErrInvalidJWT.Error())
		assert.Nil(t, newClaims)
	})
	t.Run("JWT_VALID", func(t *testing.T) {
		claims := jwt.StandardClaims{
			Audience:  "TestAudience",
			ExpiresAt: time.Now().Add(10 * time.Hour * 24).Unix(),
			Id:        "testID1234",
			IssuedAt:  time.Now().Add(-5 * time.Hour).Unix(),
			Issuer:    "TestIssuer",
			Subject:   "TestSubject",
		}
		issueJWT, err := IssueJWT(claims, "testSecret")
		assert.NoError(t, err)
		newClaims, err := Verify(user.JWT(issueJWT), "testSecret")
		assert.NoError(t, err)
		assert.Equal(t, claims.Audience, newClaims.Audience)
		assert.Equal(t, claims.ExpiresAt, newClaims.ExpiresAt)
		assert.Equal(t, claims.Id, newClaims.Id)
		assert.Equal(t, claims.IssuedAt, newClaims.IssuedAt)
		assert.Equal(t, claims.Issuer, newClaims.Issuer)
		assert.Equal(t, claims.Subject, newClaims.Subject)
	})
}
