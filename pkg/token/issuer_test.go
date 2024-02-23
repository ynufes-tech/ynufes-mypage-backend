package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"ynufes-mypage-backend/pkg/jwt"
)

func TestIssuer_IssueToken(t *testing.T) {
	jwtSecret := "jwtSecretThisShouldBe32Bytes3232"
	issuer := "testme.shion.pro"
	maxAge := 2 * time.Hour
	id1, id2 := "id1", "id2"

	tests := []struct {
		name string
		tc   func(t *testing.T)
	}{
		{
			name: "Success",
			tc: func(t *testing.T) {
				t.Parallel()
				iss := NewTokenIssuer(jwtSecret, issuer, maxAge)
				code1, err := iss.IssueNewCode(id1)
				assert.NoError(t, err)
				code2, err := iss.IssueNewCode(id2)
				assert.NoError(t, err)
				token1, err := iss.IssueToken(code1)
				assert.NoError(t, err)
				token2, err := iss.IssueToken(code2)
				assert.NoError(t, err)
				authID1, err := jwt.Verify(token1, jwtSecret)
				assert.NoError(t, err)
				authID2, err := jwt.Verify(token2, jwtSecret)
				assert.NoError(t, err)
				assert.Equal(t, id1, authID1.Id)
				assert.Equal(t, id2, authID2.Id)
			},
		},
		{
			name: "MultipleInvalidRequest",
			tc: func(t *testing.T) {
				t.Parallel()
				iss := NewTokenIssuer(jwtSecret, issuer, maxAge)
				code, err := iss.IssueNewCode(id1)
				assert.NoError(t, err)
				token1, err := iss.IssueToken(code)
				assert.NoError(t, err)
				_, err = iss.IssueToken(code)
				assert.Error(t, err)
				assert.Equal(t, "code not found", err.Error())
				tokenID1, err := jwt.Verify(token1, jwtSecret)
				assert.NoError(t, err)
				assert.Equal(t, id1, tokenID1.Id)
			},
		},
		{
			name: "ExpiredCode",
			// this testcase examines the expiration of the code
			tc: func(t *testing.T) {
				t.Parallel()
				iss := NewTokenIssuer(jwtSecret, issuer, maxAge)
				code1, err := iss.IssueNewCode(id1)
				assert.NoError(t, err)
				code2, err := iss.IssueNewCode(id2)
				assert.NoError(t, err)

				// get token1 from code1 after 4.5 seconds
				time.Sleep(4500 * time.Millisecond)
				token1, err := iss.IssueToken(code1)
				assert.NoError(t, err)
				authID1, err := jwt.Verify(token1, jwtSecret)
				assert.NoError(t, err)
				assert.Equal(t, id1, authID1.Id)

				// get token2 from code2 after 5.0 seconds
				time.Sleep(500 * time.Millisecond)
				_, err = iss.IssueToken(code2)
				assert.Error(t, err)
				assert.Equal(t, "code expired", err.Error())
			},
		},
		{
			name: "RevokedCodes",
			// scenario:
			// 1. issue code1
			// 2. wait for 0.2 seconds
			// 3. issue code2
			// 4. wait for 4.9 seconds
			// 5. issue token2 (code2 is still valid, but code1 is expired)
			// 6. wait for 0.1 seconds (wait for clean process to be done)
			// 7. issue token1 (code1 is expired, and already cleaned from the cache)
			tc: func(t *testing.T) {
				t.Parallel()
				iss := NewTokenIssuer(jwtSecret, issuer, maxAge)
				code1, err := iss.IssueNewCode(id1)
				assert.NoError(t, err)
				time.Sleep(200 * time.Millisecond)
				code2, err := iss.IssueNewCode(id2)
				assert.NoError(t, err)
				time.Sleep(4900 * time.Millisecond)

				token2, err := iss.IssueToken(code2)
				assert.NoError(t, err)
				authID2, err := jwt.Verify(token2, jwtSecret)
				assert.NoError(t, err)
				assert.Equal(t, id2, authID2.Id)
				time.Sleep(100 * time.Millisecond)
				_, err = iss.IssueToken(code1)
				assert.Error(t, err)
				assert.Equal(t, "code not found", err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.tc)
	}
}
