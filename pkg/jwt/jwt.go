package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
	"ynufes-mypage-backend/svc/pkg/exception"
)

func CreateClaims(id string, duration time.Duration, issuer string) jwt.StandardClaims {
	return jwt.StandardClaims{
		Id:        id,
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(duration * time.Second).Unix(),
		IssuedAt:  time.Now().Unix(),
	}
}

func IssueJWT(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Verify(j, secret string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(string(j), jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("UNEXPECTED SIGNING METHOD")
		}
		if token.Valid {
			return nil, exception.ErrInvalidJWT
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
