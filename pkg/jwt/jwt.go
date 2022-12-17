package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type JWT string

func IssueJWT(claims jwt.Claims, secret string) (JWT, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return JWT(tokenString), nil
}

func (j JWT) String() string {
	return string(j)
}

var ErrInvalidJWT = errors.New("INVALID JWT")

func (j JWT) Verify(secret string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(string(j), jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("UNEXPECTED SIGNING METHOD")
		}
		if token.Valid {
			return nil, ErrInvalidJWT
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
