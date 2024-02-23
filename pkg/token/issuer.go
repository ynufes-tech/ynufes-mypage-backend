package token

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"google.golang.org/appengine/v2/log"
	"strconv"
	"time"
	"ynufes-mypage-backend/pkg/jwt"
)

type Issuer struct {
	jwtSecret string
	issuer    string
	maxAge    time.Duration
	codeCache map[string]credit
}

type credit struct {
	id   string
	time int64
}

func NewTokenIssuer(jwtSecret string, issuer string, maxAge time.Duration) Issuer {
	return Issuer{
		jwtSecret: jwtSecret,
		issuer:    issuer,
		maxAge:    maxAge,
		codeCache: make(map[string]credit),
	}
}

func (v Issuer) IssueNewCode(id string) (string, error) {
	cnt := 0
	newCode, err := v.generateSecret()
	if err != nil {
		log.Warningf(context.Background(), "Failed to generate new code: %v", err)
	}
	for _, duplicate := v.codeCache[newCode]; duplicate && cnt < 10; {
		newCode, err = v.generateSecret()
		if err != nil {
			log.Warningf(context.Background(), "Failed to generate new code: %v", err)
		}
		cnt++
	}
	if cnt == 10 {
		return "", fmt.Errorf("failed to generate new code")
	}
	v.codeCache[newCode] = credit{id, time.Now().UnixMilli()}
	return newCode, nil
}

func (v Issuer) IssueToken(code string) (string, error) {
	credit, ok := v.codeCache[code]
	if !ok {
		return "", fmt.Errorf("code not found")
	}
	now := time.Now().UnixMilli()
	if now-credit.time > 5000 {
		return "", fmt.Errorf("code expired")
	}
	delete(v.codeCache, code)
	claim := jwt.CreateClaims(credit.id, v.maxAge, v.issuer)
	token, err := jwt.IssueJWT(claim, v.jwtSecret)
	if err != nil {
		return "", err
	}
	go v.RevokeOldCodes()
	return token, nil
}

func (v Issuer) RevokeOldCodes() {
	for s, t := range v.codeCache {
		// If the code is older than 5 seconds, delete it
		if time.Now().UnixMilli()-t.time > 5000 {
			delete(v.codeCache, s)
		}
	}
}

func (v Issuer) generateSecret() (string, error) {
	var secret uint64
	if err := binary.Read(rand.Reader, binary.BigEndian, &secret); err != nil {
		return "", fmt.Errorf("failed to read random number: %w", err)
	}
	return strconv.FormatUint(secret, 36), nil
}
